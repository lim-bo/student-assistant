import { useQuery } from "@tanstack/react-query";
import { globalMockData } from "@/lib/api";
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
            // Find the last in-person class (skip remote classes)
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
            const rawSchedule = globalMockData.schedule.find((item) => {
                return (
                    item.date.getDate() === day.getDate() &&
                    item.group === group
                );
            });

            if (!rawSchedule) {
                return undefined;
            }

            return transformTimeline(rawSchedule.timeline);
        },
    });
};
