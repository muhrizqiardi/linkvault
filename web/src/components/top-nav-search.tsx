import {
  HiArrowLeft, HiMagnifyingGlass
} from 'react-icons/hi2';
import { Button } from './ui/button';
import { Input } from './ui/input';

export default function TopNavSearch() {
  return (
    <header className="flex h-14 items-center px-4 border-b">
      <nav className="w-full flex gap-2">
        <Button variant="outline" size="sm" className="flex gap-4">
          <HiArrowLeft />
          <span className="sr-only">Close search</span>
        </Button>
        <Input className="h-9" />
        <Button variant="outline" size="sm" className="flex gap-4">
          <HiMagnifyingGlass />
          <span className="sr-only">Search</span>
        </Button>
      </nav>
    </header>
  );
}
