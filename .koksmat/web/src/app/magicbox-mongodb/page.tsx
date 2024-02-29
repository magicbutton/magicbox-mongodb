"use client";
import { useProcess } from "@/koksmat/useprocess";
import { APPNAME } from "../appid";

export default function Index() {
  const { isLoading, error, data } = useProcess(APPNAME, ["-h"], 20, "echo");

  return (
    <div>
      <pre>{data}</pre>
    </div>
  );
}
