import RouteCard from "./RouteCard";
import {nanoid} from "nanoid";
import {TimelineItem} from "@/lib/api";

export default function RoutesList({data}: { data: TimelineItem[] }) {
  return (
    <ol className="space-y-2 mb-4">
      {data.map((item) => (
        <RouteCard
          item={item}
          key={nanoid()}
        />
      ))}
    </ol>
  )
}