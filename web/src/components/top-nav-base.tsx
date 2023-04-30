import { HiBarsArrowDown, HiChevronDoubleLeft, HiMagnifyingGlass } from 'react-icons/hi2';
import { Button } from './ui/button';

export function TopNavBase() {
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
        <Button variant="outline" size="sm">
          <HiBarsArrowDown />
        </Button>
      </nav>
    </header>
  );
}
