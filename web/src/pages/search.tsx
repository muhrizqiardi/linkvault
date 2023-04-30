import { LinkListCard } from '@/components/link-list-card';
import { SidebarBase } from '@/components/sidebar-base';
import { TopNavBase } from '@/components/top-nav-base';
import TopNavSearch from '@/components/top-nav-search';

export default function SearchPage() {
  return (
    <>
      <div className="grid grid-cols-[256px_576px_minmax(0,1fr)]">
        <div className="">
          <SidebarBase className="border-r" />
        </div>
        <div className="border-r">
          <TopNavSearch />
          <div className="flex flex-col">
            <LinkListCard />
          </div>
        </div>
        <div className=""></div>
      </div>
    </>
  );
}
