'use client';

import { Button } from '@/components/ui/button';
import { cn } from '@/utils/cn';
import { useUser } from '@/utils/use-user';
import { DropdownMenu } from '@radix-ui/react-dropdown-menu';
import Link from 'next/link';
import {
  HiEllipsisVertical,
  HiFolder,
  HiFolderPlus,
  HiHome,
} from 'react-icons/hi2';
import SidebarBaseFolderList from './sidebar-base-folder-list';
import { Avatar, AvatarFallback } from './ui/avatar';
import {
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from './ui/dropdown-menu';
import { Skeleton } from './ui/skeleton';

interface BaseSidebarProps extends React.HTMLAttributes<HTMLDivElement> {}

export function SidebarBase({ className }: BaseSidebarProps) {
  const { data, isLoading } = useUser();

  const userButtonSkeleton = (
    <div className="flex items-center gap-4 mb-4 px-2 py-1">
      <Skeleton className="h-7 w-7 rounded-full flex-shrink-0" />
      <Skeleton className="h-4 w-full" />
    </div>
  );

  return (
    <aside className={cn('pb-12', className)}>
      <div className="space-y-4 py-4">
        <div className="px-4 py-2">
          <div className="space-y-1">
            {isLoading || data === undefined ? (
              userButtonSkeleton
            ) : (
              <DropdownMenu>
                <DropdownMenuTrigger className="w-full" asChild>
                  <Button
                    variant="ghost"
                    className="justify-start h-min mb-4 px-2 py-1"
                  >
                    <Avatar className="h-7 w-7 text-xs">
                      <AvatarFallback className="uppercase">
                        {data.full_name
                          .split(' ')
                          .map((x) => x[0].toUpperCase())
                          .join('')}
                      </AvatarFallback>
                    </Avatar>
                    <span className="ml-2">{data.full_name}</span>
                  </Button>
                </DropdownMenuTrigger>
                <DropdownMenuContent className="w-56">
                  <DropdownMenuLabel>user@example.com</DropdownMenuLabel>
                  <DropdownMenuSeparator />
                  <DropdownMenuGroup>
                    <DropdownMenuItem>
                      <span>Settings</span>
                    </DropdownMenuItem>
                    <DropdownMenuItem asChild>
                      <a href="/sign-out">Sign out</a>
                    </DropdownMenuItem>
                  </DropdownMenuGroup>
                </DropdownMenuContent>
              </DropdownMenu>
            )}
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

        <SidebarBaseFolderList />

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
