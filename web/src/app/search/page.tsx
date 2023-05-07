import { LinkListCard } from '@/components/link-list-card';
import { SidebarBase } from '@/components/sidebar-base';
import TopNavSearch from '@/components/top-nav-search';

export default function SearchPage() {
  return (
    <>
      <div className="h-14">
        <TopNavSearch />
      </div>
      <div className="flex flex-col">
        <LinkListCard />
      </div>
    </>
  );
}
