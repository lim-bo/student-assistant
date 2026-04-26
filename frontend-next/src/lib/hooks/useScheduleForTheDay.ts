import { useQuery } from "@tanstack/react-query";
import { API_CONFIG, globalMockData } from "@/lib/api";
import { nanoid } from "nanoid";

interface TransformedRemoteClass {
    id: string;
    type: "remote";
    time_start: string;
    time_end: string;
    title: string;
}

interface TransformedEvent {
    type: "event";
    time_start: string;
    time_end: string;
    title: string;
    location: string;
    coords: [number, number];
    tags: string[];
    external_url: string;
    id: string;
}

interface TransformedClass {
    id: string;
    type: "class";
    time_start: string;
    time_end: string;
    title: string;
    location: string;
    coords: [number, number];
    events: TransformedEvent[];
}

type TransformedClasses = TransformedRemoteClass | TransformedClass;

type TransformedSchedule = TransformedClasses[];

const transformTimeline = (
    timeline: Array<
        | {
              type: "class";
              time_start: string;
              time_end: string;
              title: string;
              location: string;
              coords: [number, number];
          }
        | {
              type: "event";
              time_start: string;
              time_end: string;
              title: string;
              location: string;
              coords: [number, number];
              tags: string[];
              external_url: string;
          }
        | {
              type: "remote";
              time_start: string;
              time_end: string;
              title: string;
          }
    >,
): TransformedClasses[] => {
    const result: TransformedClasses[] = [];

    for (const item of timeline) {
        if (item.type === "class") {
            result.push({
                ...item,
                id: nanoid(),
                events: [],
            });
        } else if (item.type === "remote") {
            result.push({
                ...item,
                id: nanoid(),
            });
        } else if (item.type === "event" && result.length > 0) {
            for (let i = result.length - 1; i >= 0; i--) {
                if (result[i].type === "class") {
                    (result[i] as TransformedClass).events.push({
                        ...item,
                        id: nanoid(),
                    });
                    break;
                }
            }
        }
    }

    return result;
};

export const useScheduleForTheDay = (day: Date, group: number) => {
    return useQuery<TransformedSchedule | undefined>({
        queryKey: ["schedule", day.toDateString(), group],
        queryFn: async () => {
            const date = day.toISOString().split("T")[0];
            const response = await fetch(`${API_CONFIG.BASE_URL}/schedule?date=${date}&group=${group}`);
            if (!response.ok) {
                console.log("Fetch error");
                return undefined;
            }

            const body = await response.json();
            return transformTimeline(body.timeline);
        },
    });
};
