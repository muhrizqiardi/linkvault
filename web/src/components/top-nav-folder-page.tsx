'use client';

import { FolderEntity } from '@/entities';
import Link from 'next/link';
import { useEffect, useState } from 'react';
import {
  HiBars3,
  HiBarsArrowDown,
  HiChevronDoubleLeft,
  HiMagnifyingGlass,
  HiPencil,
  HiPlus,
} from 'react-icons/hi2';
import { SidebarSheet } from './sidebar-sheet';
import { Button, buttonVariants } from './ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuLabel,
  DropdownMenuRadioGroup,
  DropdownMenuRadioItem,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from './ui/dropdown-menu';
import { Input } from './ui/input';
import { Label } from './ui/label';
import { Popover, PopoverContent, PopoverTrigger } from './ui/popover';
import { Sheet, SheetContent, SheetTrigger } from './ui/sheet';
import { Skeleton } from './ui/skeleton';

interface TopNavFolderPageProps {
  folderDetail: FolderEntity;
}

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

export function TopNavFolderPage(props: TopNavFolderPageProps) {
  const [isClientSide, setIsClientSide] = useState<boolean>(false);
  const [orderedByValue, setOrderedByValue] = useState<
    | 'title_ASC'
    | 'title_DESC'
    | 'createdAt_ASC'
    | 'createdAt_DESC'
    | 'updatedAt_ASC'
    | 'updatedAt_DESC'
    | string
  >('updatedAt_DESC');
  useEffect(() => {
    setIsClientSide(true);
  }, []);

  if (!isClientSide) return null;

  return (
    <header className="flex h-14 items-center px-4 border-b">
      <nav className="w-full flex gap-2 items-center">
        <Sheet>
          <SheetTrigger asChild>
            <Button variant="outline" size="sm" className="mr-auto lg:hidden">
              <HiBars3 />
            </Button>
          </SheetTrigger>
          <SheetContent className="w-72" position="left">
            <SidebarSheet />
          </SheetContent>
        </Sheet>

        <Button variant="link" justify="start" className="mx-auto">
          <>
            {props.folderDetail.name} <HiPencil className="ml-2" />
          </>
        </Button>

        <Link
          className={`${buttonVariants({
            variant: 'outline',
            size: 'sm',
          })} flex gap-4`}
          href="/search"
        >
          <HiMagnifyingGlass />
          <span className="sr-only">Search</span>
        </Link>
        <Popover>
          <PopoverTrigger asChild>
            <Button variant="outline" size="sm">
              <HiPlus />
            </Button>
          </PopoverTrigger>
          <PopoverContent>
            <div className="grid gap-4">
              <div className="space-y-2">
                <h4 className="font-medium leading-none">Add new link</h4>
              </div>
              <div className="grid gap-2">
                <div className="grid grid-cols-3 items-center gap-4">
                  <Label htmlFor="link-title">Title</Label>
                  <Input
                    id="link-title"
                    placeholder="Insert title here..."
                    className="col-span-2 h-8"
                  />
                </div>
                <div className="grid grid-cols-3 items-center gap-4">
                  <Label htmlFor="link-url">URL</Label>
                  <Input
                    id="link-url"
                    type="url"
                    placeholder="ex: https://twitter.com/user"
                    className="col-span-2 h-8"
                  />
                </div>

                <Button>Add</Button>
              </div>
            </div>
          </PopoverContent>
        </Popover>

        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <Button variant="outline" size="sm">
              <HiBarsArrowDown />
            </Button>
          </DropdownMenuTrigger>
          <DropdownMenuContent className="w-56">
            <DropdownMenuLabel>Order by</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuRadioGroup
              value={orderedByValue}
              onValueChange={setOrderedByValue}
            >
              <DropdownMenuRadioItem value="title_ASC">
                Title (ascending)
              </DropdownMenuRadioItem>
              <DropdownMenuRadioItem value="title_DESC">
                Title (descending)
              </DropdownMenuRadioItem>
              <DropdownMenuRadioItem value="createdAt_ASC">
                Date created (ascending)
              </DropdownMenuRadioItem>
              <DropdownMenuRadioItem value="createdAt_DESC">
                Date created (descending)
              </DropdownMenuRadioItem>
              <DropdownMenuRadioItem value="updatedAt_ASC">
                Date modified (ascending)
              </DropdownMenuRadioItem>
              <DropdownMenuRadioItem value="updatedAt_DESC">
                Date modified (descending)
              </DropdownMenuRadioItem>
            </DropdownMenuRadioGroup>
          </DropdownMenuContent>
        </DropdownMenu>
      </nav>
    </header>
  );
}
