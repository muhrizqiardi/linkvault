import { LinkListCardSkeleton } from '@/components/link-list-card-skeleton';
import { TopNavFolderPageSkeleton } from '@/components/top-nav-folder-page-skeleton';

export default function FolderPageLoading() {
  return (
    <>
      <div className="h-14">
        <TopNavFolderPageSkeleton />
      </div>
      <LinkListCardSkeleton />
      <LinkListCardSkeleton />
      <LinkListCardSkeleton />
      <LinkListCardSkeleton />
    </>
  );
}
