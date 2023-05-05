import { LinkListCard } from '@/components/link-list-card';
import { SidebarBase } from '@/components/sidebar-base';
import { TopNavBase } from '@/components/top-nav-base';

export default async function HomePage() {
  return (
    <div className="grid lg:grid-cols-[256px_576px_minmax(0,1fr)]">
      <div className="">
        <SidebarBase className="border-r hidden lg:block" />
      </div>
      <div className="border-r">
        <div className="h-14">
          <TopNavBase />
        </div>
        <div className="flex flex-col">
          <LinkListCard />
        </div>
      </div>
      <div className=""></div>
    </div>
  );
}
