import { TimelineItem } from "@/lib/api";

interface ClassItem {
    id?: string;
    type: "class";
    time_start: string;
    time_end: string;
    title: string;
    location?: string;
    coords: [number, number];
}

interface EventItem {
    id?: string;
    type: "event";
    time_start: string;
    time_end: string;
    title: string;
    location: string;
    coords: [number, number];
    tags: string[];
    external_url: string;
}

interface RemoteItem {
    id?: string;
    type: "remote";
    time_start: string;
    time_end: string;
    title: string;
}

export type RouteTimelineItem = ClassItem | EventItem | RemoteItem;

export function buildRouteTimeline(
    schedule: Array<(ClassItem | RemoteItem) & { events?: EventItem[] }>,
    selectedEvents: Record<string, EventItem>,
): TimelineItem[] {
    const result: RouteTimelineItem[] = [];

    for (const item of schedule) {
        if (item.type === "class") {
            const id = item.id ?? crypto.randomUUID();
            result.push({ ...item, id });

            if (selectedEvents[id]) {
                result.push({ ...selectedEvents[id], id: crypto.randomUUID() });
            }
        }
    }

    return result as unknown as TimelineItem[];
}
