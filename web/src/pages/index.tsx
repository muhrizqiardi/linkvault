import { LinkListCard } from '@/components/link-list-card';
import { SidebarBase } from '@/components/sidebar-base';
import { TopNavBase } from '@/components/top-nav-base';
import { GetServerSideProps } from 'next';

export const getServerSideProps: GetServerSideProps = async (context) => {
  if (context.req.cookies['token'])
    return {
      props: {},
    };

  return {
    redirect: {
      destination: '/sign-in',
    },
    props: {},
  };
};

export default function Home() {
  return (
    <>
      <div className="grid lg:grid-cols-[256px_576px_minmax(0,1fr)]">
        <div className="">
          <SidebarBase className="border-r hidden lg:block" />
        </div>
        <div className="border-r">
          <TopNavBase />
          <div className="flex flex-col">
            <LinkListCard />
          </div>
        </div>
        <div className=""></div>
      </div>
    </>
  );
}
