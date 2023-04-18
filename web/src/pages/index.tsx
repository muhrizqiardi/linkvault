import BaseLayout from '@/components/BaseLayout';
import ListViewLinkItem from '@/components/ListViewLinkItem';
import Link from 'next/link';
import { HiBars3, HiMagnifyingGlass, HiPlus } from 'react-icons/hi2';

export default function Home() {
  return (
    <BaseLayout>
      <div className="flex-shrink-0">
        <nav className="navbar border-b border-b-base-200 gap-2">
          <label
            htmlFor="main-drawer"
            className="btn btn-ghost btn-square drawer-button lg:hidden text-2xl"
          >
            <HiBars3 />
          </label>
          <h1 className="mr-auto lg:ml-2">All folders</h1>
          <Link href="#" className="btn btn-ghost btn-square text-2xl">
            <HiMagnifyingGlass />
          </Link>
          <Link href="#" className="btn btn-ghost btn-square text-2xl">
            <HiPlus />
          </Link>
        </nav>
        <div className="flex flex-col md:mt-4 md:w-full md:max-w-6xl md:mx-auto border border-base-200 overflow-hidden">
          <ListViewLinkItem />
        </div>
      </div>
    </BaseLayout>
  );
}
