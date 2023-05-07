import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from '@/components/ui/card';
import { LinkEntity } from '@/entities';
import Link from 'next/link';
import { HiFolder } from 'react-icons/hi2';
import {
  ContextMenu,
  ContextMenuContent,
  ContextMenuItem,
  ContextMenuTrigger,
} from './ui/context-menu';

interface LinkListCardProps {
  link: LinkEntity;
}

export function LinkListCard(props: LinkListCardProps) {
  return (
    <ContextMenu>
      <ContextMenuTrigger>
        <Link href="/">
          <Card className="border-t-0 border-r-0 border-l-0 shadow-none rounded-none hover:bg-black hover:bg-opacity-5">
            <div className="flex">
              <div className="p-3.5 pr-0 flex-shrink-0">
                {props.link?.cover_url !== undefined ? (
                  <img
                    src={props.link.cover_url}
                    alt={props.link.title}
                    className="aspect-[1200/680] w-24 object-contain bg-neutral-200 rounded-sm"
                  />
                ) : (
                  <div className="aspect-[1200/680] w-24 bg-neutral-200 rounded-sm"></div>
                )}
              </div>
              <div className="block">
                <CardHeader>
                  <CardDescription className="inline-flex items-center">
                    {new URL(props.link.url).hostname} · 2h ago ·{' '}
                    <span className="ml-2 inline-flex items-center">
                      <HiFolder className="mr-1" />
                      Bookmarked Tweets
                    </span>
                  </CardDescription>
                  <CardTitle>{props.link.title}</CardTitle>
                </CardHeader>
                <CardContent className="text-sm">
                  {props.link.excerpt}
                </CardContent>
              </div>
            </div>
          </Card>
        </Link>
      </ContextMenuTrigger>
      <ContextMenuContent className="w-64">
        <ContextMenuItem inset>Edit</ContextMenuItem>
        <ContextMenuItem inset>Select</ContextMenuItem>
        <ContextMenuItem inset>Copy full link address</ContextMenuItem>
        <ContextMenuItem inset>Open in new tab</ContextMenuItem>
      </ContextMenuContent>
    </ContextMenu>
  );
}
