'use client';

import {
  CreateNewLinkDto,
  createNewLinkDtoSchema,
  FolderEntity,
} from '@/entities';
import { zodResolver } from '@hookform/resolvers/zod';
import { Loader2 } from 'lucide-react';
import { useRouter } from 'next/navigation';
import { useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { HiPlus } from 'react-icons/hi2';
import { Button } from './ui/button';
import { Input } from './ui/input';
import { Label } from './ui/label';
import { Popover, PopoverContent, PopoverTrigger } from './ui/popover';
import { Textarea } from './ui/textarea';

interface TopNavFolderPageNewLinkPopupProps {
  folderDetail: FolderEntity;
}

export function TopNavFolderPageNewLinkPopup(
  props: TopNavFolderPageNewLinkPopupProps,
) {
  const router = useRouter();
  const [popoverIsOpen, setPopoverIsOppen] = useState(false);
  const [isLoading, setIsLoading] = useState(false);
  const { register, handleSubmit } = useForm<CreateNewLinkDto>({
    resolver: zodResolver(createNewLinkDtoSchema),
  });

  const onSubmit: SubmitHandler<CreateNewLinkDto> = async (data) => {
    setIsLoading(true);
    try {
      const response = await fetch(
        `/api/folders/${props.folderDetail.id}/links`,
        {
          method: 'POST',
          body: JSON.stringify(data),
        },
      );
      if (!response.ok) throw response;
      setPopoverIsOppen(false);
      router.refresh();
    } catch (error) {
      console.error(error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Popover open={popoverIsOpen} onOpenChange={setPopoverIsOppen}>
      <PopoverTrigger asChild>
        <Button variant="outline" size="sm">
          <HiPlus />
        </Button>
      </PopoverTrigger>
      <PopoverContent>
        <form onSubmit={handleSubmit(onSubmit)} className="grid gap-4">
          <div className="space-y-2">
            <h4 className="font-medium leading-none">
              Add new link on the folder "{props.folderDetail.name}"
            </h4>
          </div>
          <img
            src="https://picsum.photos/id/237/200/300"
            className="aspect-[1200/628] w-full object-contain bg-gray-100 rounded-lg"
          />
          <div className="grid gap-2">
            <div className="grid grid-cols-3 items-center gap-4">
              <Label htmlFor="link-cover-url">Cover URL</Label>
              <Input
                id="link-cover-url"
                type="url"
                placeholder="ex: https://twitter.com/user"
                className="col-span-2 h-8"
                {...register('cover_url')}
                required
              />
            </div>
            <div className="grid grid-cols-3 items-center gap-4">
              <Label htmlFor="link-title">Title</Label>
              <Input
                id="link-title"
                placeholder="Insert title here..."
                className="col-span-2 h-8"
                {...register('title')}
                required
              />
            </div>
            <div className="grid grid-cols-3 items-center gap-4">
              <Label htmlFor="link-url">URL</Label>
              <Input
                id="link-url"
                type="url"
                placeholder="ex: https://twitter.com/user"
                className="col-span-2 h-8"
                {...register('url')}
                required
              />
            </div>
            <div className="grid grid-cols-3 items-center gap-4">
              <Label htmlFor="link-excerpt">Excerpt</Label>
              <Textarea
                id="link-excerpt"
                className="col-span-2"
                {...register('excerpt')}
                required
              />
            </div>
            <Button disabled={isLoading} type="submit" className="mt-4">
              {isLoading ? (
                <Loader2 className="mr-2 h-4 w-4 animate-spin" />
              ) : null}
              Save
            </Button>
          </div>
        </form>
      </PopoverContent>
    </Popover>
  );
}
