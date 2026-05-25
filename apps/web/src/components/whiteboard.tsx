"use client";

import "@excalidraw/excalidraw/index.css";

import dynamic from "next/dynamic";

const Excalidraw = dynamic(
  async () => {
    const mod = await import(
      "@excalidraw/excalidraw"
    );

    return mod.Excalidraw;
  },
  {
    ssr: false,

    loading: () => (
      <div className="p-6">
        Loading Excalidraw...
      </div>
    ),
  }
);

export default function Whiteboard() {
  return (
    <div className="h-screen w-full">
      <Excalidraw />
    </div>
  );
}