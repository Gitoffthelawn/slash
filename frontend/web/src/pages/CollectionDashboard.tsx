import { Button, Input } from "@mui/joy";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import CollectionView from "@/components/CollectionView";
import CreateCollectionDrawer from "@/components/CreateCollectionDrawer";
import { shortcutService } from "@/services";
import useCollectionStore from "@/stores/v1/collection";
import FilterView from "../components/FilterView";
import Icon from "../components/Icon";
import useLoading from "../hooks/useLoading";

interface State {
  showCreateCollectionDrawer: boolean;
}

const CollectionDashboard: React.FC = () => {
  const { t } = useTranslation();
  const loadingState = useLoading();
  const collectionStore = useCollectionStore();
  const [state, setState] = useState<State>({
    showCreateCollectionDrawer: false,
  });
  const [search, setSearch] = useState<string>("");
  const filteredCollections = collectionStore.getCollectionList().filter((collection) => {
    return (
      collection.name.toLowerCase().includes(search.toLowerCase()) ||
      collection.title.toLowerCase().includes(search.toLowerCase()) ||
      collection.description.toLowerCase().includes(search.toLowerCase())
    );
  });

  useEffect(() => {
    Promise.all([shortcutService.getMyAllShortcuts(), collectionStore.fetchCollectionList()]).finally(() => {
      loadingState.setFinish();
    });
  }, []);

  const setShowCreateCollectionDrawer = (show: boolean) => {
    setState({
      ...state,
      showCreateCollectionDrawer: show,
    });
  };

  return (
    <>
      <div className="mx-auto max-w-8xl w-full px-3 md:px-12 pt-4 pb-6 flex flex-col justify-start items-start">
        <div className="w-full flex flex-row justify-start items-start mb-4">
          <div className="bg-yellow-100 dark:bg-yellow-500 dark:opacity-70 py-2 px-3 rounded-full border dark:border-yellow-600 flex flex-row justify-start items-center cursor-pointer shadow">
            <Icon.LibrarySquare className="w-5 h-auto opacity-60" />
            <a
              className="hover:underline hover:text-blue-600"
              href="https://github.com/boojack/slash/blob/main/docs/getting-started/collections.md"
              target="_blank"
            >
              <span className="mx-1 text-sm">Collection is in Beta. Learn more in docs</span>
              <Icon.ExternalLink className="w-4 h-auto inline-block" />
            </a>
          </div>
        </div>
        <div className="w-full flex flex-row justify-between items-center mb-4">
          <div>
            <Input
              className="w-32 mr-2"
              type="text"
              size="sm"
              placeholder={t("common.search")}
              startDecorator={<Icon.Search className="w-4 h-auto" />}
              value={search}
              onChange={(e) => setSearch(e.target.value)}
            />
          </div>
          <div className="flex flex-row justify-start items-center">
            <Button className="hover:shadow" variant="soft" size="sm" onClick={() => setShowCreateCollectionDrawer(true)}>
              <Icon.Plus className="w-5 h-auto" />
              <span className="ml-0.5">{t("common.create")}</span>
            </Button>
          </div>
        </div>
        <FilterView />
        {loadingState.isLoading ? (
          <div className="py-12 w-full flex flex-row justify-center items-center opacity-80 dark:text-gray-500">
            <Icon.Loader className="mr-2 w-5 h-auto animate-spin" />
            {t("common.loading")}
          </div>
        ) : filteredCollections.length === 0 ? (
          <div className="py-16 w-full flex flex-col justify-center items-center text-gray-400">
            <Icon.PackageOpen className="w-16 h-auto" strokeWidth="1" />
            <p className="mt-4">No collections found.</p>
          </div>
        ) : (
          <div className="w-full flex flex-col justify-start items-start gap-3">
            {filteredCollections.map((collection) => {
              return <CollectionView key={collection.id} collection={collection} />;
            })}
          </div>
        )}
      </div>

      {state.showCreateCollectionDrawer && (
        <CreateCollectionDrawer
          onClose={() => setShowCreateCollectionDrawer(false)}
          onConfirm={() => setShowCreateCollectionDrawer(false)}
        />
      )}
    </>
  );
};

export default CollectionDashboard;