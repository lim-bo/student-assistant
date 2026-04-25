import { useMutation } from "@tanstack/react-query";

interface EventItem {
    id: string;
    type: "event";
    time_start: string;
    time_end: string;
    title: string;
    location: string;
    coords: [number, number];
    tags: string[];
    external_url: string;
}

interface RouteTimelineItem {
    type: "class" | "event";
    coords: [number, number];
}

type RoutePayload = RouteTimelineItem[];

interface SubmitRouteParams {
    schedule: Array<{
        id: string;
        type: "class" | "remote";
        time_start: string;
        time_end: string;
        title: string;
        location?: string;
        coords?: [number, number];
        events?: EventItem[];
    }>;
    selectedEvents: Record<string, EventItem>;
}

export const useSubmitRoute = () => {
    return useMutation({
        mutationFn: async ({ schedule, selectedEvents }: SubmitRouteParams) => {
            const routeTimeline: RouteTimelineItem[] = [];

            for (const item of schedule) {
                if (item.type === "class") {
                    routeTimeline.push({
                        type: "class",
                        coords: item.coords!,
                    });

                    if (selectedEvents[item.id]) {
                        routeTimeline.push({
                            type: "event",
                            coords: selectedEvents[item.id].coords,
                        });
                    }
                }
            }

            const payload: RoutePayload = routeTimeline;

            console.log("Submitting route:", payload);

            const response = await fetch("/api/route", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                body: JSON.stringify(payload),
            });

            if (!response.ok) {
                throw new Error("Failed to submit route");
            }

            return response.json();
        },
    });
};
