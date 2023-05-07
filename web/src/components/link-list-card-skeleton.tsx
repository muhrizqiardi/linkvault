import { Skeleton } from './ui/skeleton';

export function LinkListCardSkeleton() {
  return (
    <div className="border-b flex">
      <div className="p-3.5 pr-0 flex-shrink-0">
        <Skeleton className="aspect-[1200/680] w-24" />
      </div>
      <div className="p-3.5 w-full">
        <Skeleton className="h-3 mb-2 w-1/3" />
        <Skeleton className="h-6 mb-3" />
        <Skeleton className="h-3 mb-2" />
        <Skeleton className="h-3 mb-2" />
        <Skeleton className="h-3 mb-2" />
        <Skeleton className="h-3 mb-2" />
        <Skeleton className="h-3 mb-2" />
      </div>
    </div>
  );
}
