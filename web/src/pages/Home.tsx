import { Button, Tab, TabList, Tabs } from "@mui/joy";
import { useEffect, useState } from "react";
import { shortcutService } from "../services";
import { useAppSelector } from "../stores";
import useFilterStore, { Filter } from "../stores/v1/filter";
import useUserStore from "../stores/v1/user";
import useLoading from "../hooks/useLoading";
import Icon from "../components/Icon";
import ShortcutListView from "../components/ShortcutListView";
import CreateShortcutDialog from "../components/CreateShortcutDialog";
import FilterView from "../components/FilterView";

interface State {
  showCreateShortcutDialog: boolean;
}

const getFilteredShortcutList = (shortcutList: Shortcut[], filter: Filter, currentUser: User) => {
  const { tag, mineOnly } = filter;
  const filteredShortcutList = shortcutList.filter((shortcut) => {
    if (tag) {
      if (!shortcut.tags.includes(tag)) {
        return false;
      }
    }
    if (mineOnly) {
      if (shortcut.creatorId !== currentUser.id) {
        return false;
      }
    }
    return true;
  });
  return filteredShortcutList;
};

const Home: React.FC = () => {
  const loadingState = useLoading();
  const currentUser = useUserStore().getCurrentUser();
  const filterStore = useFilterStore();
  const { shortcutList } = useAppSelector((state) => state.shortcut);
  const [state, setState] = useState<State>({
    showCreateShortcutDialog: false,
  });
  const filteredShortcutList = getFilteredShortcutList(shortcutList, filterStore.filter, currentUser);

  useEffect(() => {
    Promise.all([shortcutService.getMyAllShortcuts()]).finally(() => {
      loadingState.setFinish();
    });
  }, []);

  const setShowCreateShortcutDialog = (show: boolean) => {
    setState({
      ...state,
      showCreateShortcutDialog: show,
    });
  };

  return (
    <>
      <div className="mx-auto max-w-4xl w-full px-3 py-6 flex flex-col justify-start items-start">
        <div className="w-full flex flex-row justify-start items-center mb-4">
          <span className="font-mono text-gray-400 mr-2">Shortcuts</span>
        </div>
        <div className="w-full flex flex-row justify-between items-center mb-4">
          <div className="flex flex-row justify-start items-center">
            <Tabs defaultValue={"ALL"} size="sm" onChange={(_, value) => filterStore.setFilter({ mineOnly: value !== "ALL" })}>
              <TabList>
                <Tab value={"ALL"}>All</Tab>
                <Tab value={"PRIVATE"}>Mine</Tab>
              </TabList>
            </Tabs>
          </div>
          <div>
            <Button variant="soft" size="sm" onClick={() => setShowCreateShortcutDialog(true)}>
              <Icon.Plus className="w-5 h-auto" /> New
            </Button>
          </div>
        </div>

        <FilterView />

        {loadingState.isLoading ? (
          <div className="py-12 w-full flex flex-row justify-center items-center opacity-80">
            <Icon.Loader className="mr-2 w-5 h-auto animate-spin" />
            loading
          </div>
        ) : filteredShortcutList.length === 0 ? (
          <div className="py-4 w-full flex flex-col justify-center items-center">
            <Icon.PackageOpen className="w-12 h-auto text-gray-400" />
            <p className="mt-4 mb-2">No shortcuts found.</p>
            <Button size="sm" onClick={() => setShowCreateShortcutDialog(true)}>
              <Icon.Plus className="w-5 h-auto mr-1" />
              Create
            </Button>
          </div>
        ) : (
          <ShortcutListView shortcutList={filteredShortcutList} />
        )}
      </div>

      {state.showCreateShortcutDialog && (
        <CreateShortcutDialog onClose={() => setShowCreateShortcutDialog(false)} onConfirm={() => setShowCreateShortcutDialog(false)} />
      )}
    </>
  );
};

export default Home;
