package v1

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/fieldmaskpb"

	v1pb "github.com/yourselfhosted/slash/proto/gen/api/v1"
	"github.com/yourselfhosted/slash/store"
	teststore "github.com/yourselfhosted/slash/store/test"
)

func TestUpdateUserPasswordChangesSignInCredentials(t *testing.T) {
	ctx := context.Background()
	testingStore := teststore.NewTestingStore(ctx, t)
	t.Cleanup(func() {
		require.NoError(t, testingStore.Close())
	})

	oldPassword := "old-password"
	oldPasswordHash, err := bcrypt.GenerateFromPassword([]byte(oldPassword), bcrypt.DefaultCost)
	require.NoError(t, err)
	user, err := testingStore.CreateUser(ctx, &store.User{
		Email:        "test@example.com",
		Nickname:     "Test User",
		PasswordHash: string(oldPasswordHash),
		Role:         store.RoleUser,
	})
	require.NoError(t, err)

	service := &APIV1Service{
		Secret: "test-secret",
		Store:  testingStore,
	}
	authenticatedCtx := context.WithValue(ctx, userIDContextKey, user.ID)
	newPassword := "new-password"
	_, err = service.UpdateUser(authenticatedCtx, &v1pb.UpdateUserRequest{
		User: &v1pb.User{
			Id:       user.ID,
			Password: newPassword,
		},
		UpdateMask: &fieldmaskpb.FieldMask{Paths: []string{"password"}},
	})
	require.NoError(t, err)

	_, err = service.SignIn(ctx, &v1pb.SignInRequest{
		Email:    user.Email,
		Password: oldPassword,
	})
	require.Equal(t, codes.InvalidArgument, status.Code(err))

	transportCtx := grpc.NewContextWithServerTransportStream(ctx, noopServerTransportStream{})
	_, err = service.SignIn(transportCtx, &v1pb.SignInRequest{
		Email:    user.Email,
		Password: newPassword,
	})
	require.NoError(t, err)
}

type noopServerTransportStream struct{}

func (noopServerTransportStream) Method() string               { return "" }
func (noopServerTransportStream) SetHeader(metadata.MD) error  { return nil }
func (noopServerTransportStream) SendHeader(metadata.MD) error { return nil }
func (noopServerTransportStream) SetTrailer(metadata.MD) error { return nil }
