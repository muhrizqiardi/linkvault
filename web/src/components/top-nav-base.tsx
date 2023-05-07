'use client';

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

export function TopNavBase() {
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
      <nav className="w-full flex gap-2">
        <Button variant="outline" size="sm" className="mr-auto hidden lg:block">
          <HiChevronDoubleLeft />
        </Button>
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

        <Button variant="link" className="mr-auto">
          Folder name <HiPencil className="ml-2" />
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
