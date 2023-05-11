'use client';

import { useUser } from '@/utils/use-user';
import { Avatar, AvatarFallback } from './ui/avatar';
import { Button } from './ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuGroup,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from './ui/dropdown-menu';
import { Skeleton } from './ui/skeleton';

export function SidebarBaseUserDropdownMenu() {
  const { data, isLoading } = useUser();

  const userButtonSkeleton = (
    <div className="flex items-center gap-4 mb-4 px-2 py-1">
      <Skeleton className="h-7 w-7 rounded-full flex-shrink-0" />
      <Skeleton className="h-4 w-full" />
    </div>
  );

  return isLoading || data === undefined ? (
    userButtonSkeleton
  ) : (
    <DropdownMenu>
      <DropdownMenuTrigger className="w-full" asChild>
        <Button variant="ghost" className="justify-start h-min mb-4 px-2 py-1">
          <Avatar className="h-7 w-7 text-xs">
            <AvatarFallback className="uppercase">
              {data.full_name
                .split(' ')
                .map((x) => x[0].toUpperCase())
                .join('')
                .slice(0, 2)}
            </AvatarFallback>
          </Avatar>
          <span className="ml-2 text-left">{data.full_name}</span>
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
  );
}
