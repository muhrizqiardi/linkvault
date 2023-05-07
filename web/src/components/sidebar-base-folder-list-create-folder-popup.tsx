'use client';

import { CreateFolderPopupFormDto, createFolderPopupFormDto } from '@/schemas';
import { zodResolver } from '@hookform/resolvers/zod';
import { DialogClose } from '@radix-ui/react-dialog';
import { Loader2 } from 'lucide-react';
import { useRouter } from 'next/navigation';
import { useState } from 'react';
import { SubmitHandler, useForm } from 'react-hook-form';
import { HiFolderPlus } from 'react-icons/hi2';
import { Button } from './ui/button';
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from './ui/dialog';
import { Input } from './ui/input';
import { Label } from './ui/label';

export function SidebarBaseFolderListCreateFolderPopup() {
  const router = useRouter();
  const { handleSubmit, register } = useForm<CreateFolderPopupFormDto>({
    resolver: zodResolver(createFolderPopupFormDto),
  });
  const [isLoading, setIsLoading] = useState(false);
  const [dialogIsOpen, setDialogIsOpen] = useState(false);

  const onSubmit: SubmitHandler<CreateFolderPopupFormDto> = async (data) => {
    setIsLoading(true);
    try {
      const response = await fetch('/api/folders', {
        method: 'POST',
        body: JSON.stringify(data),
      });
      if (!response.ok) throw new Error('Request failed');
      setDialogIsOpen(false);
      router.refresh();
    } catch (error) {
      console.error(error);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <Dialog open={dialogIsOpen} onOpenChange={setDialogIsOpen}>
      <DialogTrigger asChild>
        <Button variant="ghost" size="sm" className="rounded-full h-9 w-9">
          <HiFolderPlus />
        </Button>
      </DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <form onSubmit={handleSubmit(onSubmit)}>
          <DialogHeader>
            <DialogTitle>Create New Folder</DialogTitle>
          </DialogHeader>
          <div className="grid gap-4 py-4">
            <div className="grid grid-cols-4 items-center gap-4">
              <Label htmlFor="name" className="text-right">
                Name
              </Label>
              <Input id="name" className="col-span-3" {...register('name')} />
            </div>
          </div>
          <DialogFooter>
            <Button disabled={isLoading} type="submit">
              {isLoading ? (
                <Loader2 className="mr-2 h-4 w-4 animate-spin" />
              ) : null}
              Create
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
