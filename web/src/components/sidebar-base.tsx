import { Button } from '@/components/ui/button';
import { cn } from '@/utils/cn';
import { Suspense } from 'react';
import { HiEllipsisVertical, HiFolder, HiHome } from 'react-icons/hi2';
import SidebarBaseFolderList, {
  SidebarBaseFolderListSkeleton,
} from './sidebar-base-folder-list';
import { SidebarBaseUserDropdownMenu } from './sidebar-base-user-dropdown-menu';

interface SidebarBaseProps extends React.HTMLAttributes<HTMLDivElement> {}

export function SidebarBase({ className }: SidebarBaseProps) {
  return (
    <aside className={cn('pb-12', className)}>
      <div className="space-y-4 py-4">
        <div className="px-4 py-2">
          <div className="space-y-1">
            <SidebarBaseUserDropdownMenu />
            <Button
              variant="secondary"
              size="sm"
              className="w-full justify-start"
            >
              <HiHome className="mr-2 h-4 w-4" />
              All links
            </Button>
          </div>
        </div>

        <Suspense fallback={<SidebarBaseFolderListSkeleton />}>
          <SidebarBaseFolderList />
        </Suspense>

        <div className="px-4 py-2">
          <div className="flex justify-between items-center">
            <h2 className="px-2 text-lg font-semibold tracking-tight">Tags</h2>
            <div className="flex gap-1">
              <Button variant="outline" size="sm">
                <HiEllipsisVertical />
              </Button>
            </div>
          </div>
          <div className="space-y-1 py-2">
            <Button variant="ghost" size="sm" className="w-full justify-start">
              <HiFolder className="mr-2 h-4 w-4" />
              Tweets
            </Button>
            <Button variant="ghost" size="sm" className="w-full justify-start">
              <HiFolder className="mr-2 h-4 w-4" />
              Posts
            </Button>
            <Button variant="ghost" size="sm" className="w-full justify-start">
              <HiFolder className="mr-2 h-4 w-4" />
              All about music
            </Button>
            <Button variant="ghost" size="sm" className="w-full justify-start">
              <HiFolder className="mr-2 h-4 w-4" />
              Entertainments
            </Button>
          </div>
        </div>
      </div>
    </aside>
  );
}
