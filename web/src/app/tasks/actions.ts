"use server";

import { refresh } from "next/cache";
import { baseApiUrl } from "./page";

export async function markTaskAsDone(taskID: string) {
  // @TODO: handle not ok status response
  await fetch(`${baseApiUrl}/api/tasks/${taskID}/status/done`, {
    method: "PATCH",
  });

  refresh();
}

export async function markTaskAsTodo(taskID: string) {
  // @TODO: handle not ok status response
  await fetch(`${baseApiUrl}/api/tasks/${taskID}/status/todo`, {
    method: "PATCH",
  });

  refresh();
}
