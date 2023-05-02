import { cn } from '@/utils/cn';
import {
  HiEllipsisVertical,
  HiFolder,
  HiFolderPlus,
  HiHome,
  HiUser,
} from 'react-icons/hi2';
import { Button } from './ui/button';

interface SidebarSheetProps extends React.HTMLAttributes<HTMLDivElement> {}

export function SidebarSheet(props: SidebarSheetProps) {
  return (
    <aside className={cn('pb-12', props.className)}>
      <div className="space-y-4 py-4">
        <div className="px-4 py-2">
          <div className="space-y-1">
            <Button
              variant="secondary"
              size="sm"
              className="w-full justify-start"
            >
              <HiHome className="mr-2 h-4 w-4" />
              All links
            </Button>
            <Button variant="ghost" size="sm" className="w-full justify-start">
              <HiUser className="mr-2 h-4 w-4" />
              Account
            </Button>
          </div>
        </div>

        <div className="px-4 py-2">
          <div className="flex justify-between items-center">
            <h2 className="px-2 text-lg font-semibold tracking-tight">
              Folders
            </h2>
            <div className="flex gap-1">
              <Button variant="outline" size="sm">
                <HiEllipsisVertical />
              </Button>
              <Button variant="outline" size="sm">
                <HiFolderPlus />
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
