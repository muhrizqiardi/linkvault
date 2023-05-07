'use client';

import { FolderEntity } from '@/entities';
import { HiCheck, HiFolder, HiXMark } from 'react-icons/hi2';
import { Button, buttonVariants } from './ui/button';
import { LinkWithActiveStyle } from './link-with-active-style';
import {
  ContextMenu,
  ContextMenuTrigger,
  ContextMenuContent,
  ContextMenuItem,
} from './ui/context-menu';
import Link from 'next/link';
import { useState } from 'react';
import { Input } from './ui/input';
import { SubmitHandler, useForm } from 'react-hook-form';
import { zodResolver } from '@hookform/resolvers/zod';
import { UpdateFolderDto, updateFolderDtoSchema } from '@/schemas';
import { useRouter } from 'next/navigation';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from './ui/dialog';
import { DialogClose } from '@radix-ui/react-dialog';

export function SidebarBaseFolderListItem(props: { folder: FolderEntity }) {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState(false);
  const [confirmationDialogIsOpen, setConfirmationDialogIsOpen] =
    useState(false);
  const [renameIsEnabled, setRenameIsEnabled] = useState(false);
  const { register, handleSubmit } = useForm<UpdateFolderDto>({
    resolver: zodResolver(updateFolderDtoSchema),
  });

  const handleDeleteFolder = async () => {
    setIsLoading(true);
    try {
      const response = await fetch(`/api/folders/${props.folder.id}`, {
        method: 'DELETE',
      });

      if (!response.ok) throw response;

      setRenameIsEnabled(false);
      router.refresh();
    } catch (error) {
      console.error(error);
    } finally {
      setIsLoading(false);
    }
  };

  const onSubmit: SubmitHandler<UpdateFolderDto> = async (data) => {
    setIsLoading(true);
    try {
      const response = await fetch(`/api/folders/${props.folder.id}`, {
        method: 'PATCH',
        body: JSON.stringify(data),
      });

      if (!response.ok) throw response;

      setConfirmationDialogIsOpen(false);
      router.refresh();
    } catch (error) {
      console.error(error);
    } finally {
      setIsLoading(false);
    }
  };

  if (renameIsEnabled)
    return (
      <form onSubmit={handleSubmit(onSubmit)} className="flex gap-1">
        <Input
          className="h-9"
          defaultValue={props.folder.name}
          disabled={isLoading}
          {...register('name')}
        />
        <Button
          type="button"
          onClick={() => setRenameIsEnabled(false)}
          variant="outline"
          size="sm"
        >
          <HiXMark />
        </Button>
        <Button disabled={isLoading} size="sm">
          <HiCheck />
        </Button>
      </form>
    );

  return (
    <Dialog>
      <ContextMenu>
        <ContextMenuTrigger asChild>
          <LinkWithActiveStyle
            href={`/folders/${props.folder.id}`}
            className={
              buttonVariants({
                variant: 'ghost',
                size: 'sm',
                justify: 'start',
              }) + ' w-full'
            }
            activeClassName={
              buttonVariants({
                variant: 'secondary',
                size: 'sm',
                justify: 'start',
              }) + ' w-full'
            }
          >
            <HiFolder className="mr-2 h-4 w-4" />
            {props.folder.name}
          </LinkWithActiveStyle>
        </ContextMenuTrigger>
        <ContextMenuContent>
          <ContextMenuItem onClick={() => setRenameIsEnabled((x) => !x)}>
            Rename
          </ContextMenuItem>
          <ContextMenuItem asChild>
            <DialogTrigger className="text-destructive w-full">
              Delete
            </DialogTrigger>
          </ContextMenuItem>
        </ContextMenuContent>
      </ContextMenu>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>
            Are you sure you want to delete the folder "{props.folder.name}"?
          </DialogTitle>
          <DialogDescription>This action is irreversible.</DialogDescription>
        </DialogHeader>
        <DialogFooter>
          <DialogClose asChild>
            <Button variant="destructive" onClick={handleDeleteFolder}>
              Confirm
            </Button>
          </DialogClose>
        </DialogFooter>
      </DialogContent>
    </Dialog>
  );
}
