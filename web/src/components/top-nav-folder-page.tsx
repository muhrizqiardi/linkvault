'use client';

import { FolderEntity } from '@/entities';
import { PopoverClose } from '@radix-ui/react-popover';
import Link from 'next/link';
import { useEffect, useState } from 'react';
import {
  HiBars3,
  HiBarsArrowDown,
  HiFolder,
  HiMagnifyingGlass,
  HiPencil,
  HiPlus,
} from 'react-icons/hi2';
import { SidebarSheet } from './sidebar-sheet';
import { TopNavFolderPageNewLinkPopup } from './top-nav-folder-page-new-link-popup';
import { Button, buttonVariants } from './ui/button';
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
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

interface TopNavFolderPageProps {
  folderDetail: FolderEntity;
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

        <TopNavFolderPageNewLinkPopup folderDetail={props.folderDetail} />

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
