import { TopNavFolderPageSkeleton } from '@/components/top-nav-folder-page';

export default function FolderPageLoading() {
  return (
    <>
      <div className="h-14">
        <TopNavFolderPageSkeleton />
      </div>
    </>
  );
}
