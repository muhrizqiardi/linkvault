import { useEffect, useState } from 'react';
import {
  HiBarsArrowDown,
  HiChevronDoubleLeft,
  HiMagnifyingGlass,
  HiPlus,
} from 'react-icons/hi2';
import { Button } from './ui/button';
import { Input } from './ui/input';
import { Label } from './ui/label';
import { Popover, PopoverContent, PopoverTrigger } from './ui/popover';

export function TopNavBase() {
  const [isClientSide, setIsClientSide] = useState<boolean>(false);
  useEffect(() => {
    setIsClientSide(true);
  }, []);

  if (!isClientSide) return null;

  return (
    <header className="flex h-14 items-center px-4 border-b">
      <nav className="w-full flex gap-2">
        <Button variant="outline" size="sm" className="mr-auto">
          <HiChevronDoubleLeft />
        </Button>
        <Button variant="outline" size="sm" className="flex gap-4">
          <HiMagnifyingGlass />
          <span className="sr-only">Search</span>
        </Button>
        <Popover>
          <PopoverTrigger>
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
        <Button variant="outline" size="sm">
          <HiBarsArrowDown />
        </Button>
      </nav>
    </header>
  );
}
