import { Skeleton } from './ui/skeleton';

export function TopNavFolderPageSkeleton() {
  return (
    <div className="flex h-14 items-center px-4 border-b">
      <div className="w-full flex items-center gap-2">
        <Skeleton className="h-4 w-48 mx-auto" />
        <Skeleton className="h-9 w-9" />
        <Skeleton className="h-9 w-9" />
        <Skeleton className="h-9 w-9" />
      </div>
    </div>
  );
}
