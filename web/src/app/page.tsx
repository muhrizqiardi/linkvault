import { LinkListCard } from '@/components/link-list-card';
import { SidebarBase } from '@/components/sidebar-base';
import { TopNavBase } from '@/components/top-nav-base';

export default async function HomePage() {
  return (
    <>
      <div className="h-14">
        <TopNavBase />
      </div>
      <div className="flex flex-col">
        <LinkListCard />
      </div>
    </>
  );
}
