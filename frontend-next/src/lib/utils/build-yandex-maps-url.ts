import { EventItem, InPersonClass, TimelineItem } from "@/lib/api";

export function buildRoute(items: TimelineItem[]) {
    return items
        .filter(
            (item): item is InPersonClass | EventItem => item.type !== "remote",
        )
        .map((item) => {
            return item.coords.join(",");
        })
        .join("~");
}
