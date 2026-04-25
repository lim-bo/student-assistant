interface InPersonClass {
    type: "class";
    time_start: string;
    time_end: string;
    title: string;
    location: string;
    coords: [number, number];
}

interface RemoteClass {
    type: "remote";
    time_start: string;
    time_end: string;
    title: string;
}

interface EventItem {
    type: "event";
    time_start: string;
    time_end: string;
    title: string;
    location: string;
    coords: [number, number];
    tags: string[];
    external_url: string;
}

type TimelineItem = InPersonClass | RemoteClass | EventItem;

interface Schedule {
    group: number;
    date: Date;
    timeline: TimelineItem[];
}

const schedule: Schedule[] = [
    {
        group: 3376,
        date: new Date(),
        timeline: [
            {
                type: "class",
                time_start: "11:00",
                time_end: "13:00",
                title: "Социология",
                location: "ул. F, 9, ауд. 101",
                coords: [12.0, 12.0],
            },
            {
                type: "event",
                time_start: "13:20",
                time_end: "13:40",
                title: "Лекция по Python",
                location: "ул. V, 9АК, этаж 2",
                coords: [12.0, 12.0],
                tags: ["IT", "backend"],
                external_url: "https://kudago.com/event/1231231231",
            },
            {
                type: "class",
                time_start: "14:00",
                time_end: "15:00",
                title: "Математический анализ",
                location: "ул. F, 9, ауд. 106",
                coords: [12.0, 12.0],
            },
            {
                type: "event",
                time_start: "16:00",
                time_end: "17:00",
                title: "Лекция по Python",
                location: "ул. V, 9АК, этаж 2",
                coords: [12.0, 12.0],
                tags: ["IT", "backend"],
                external_url: "https://kudago.com/event/1231231231",
            },
            {
                type: "event",
                time_start: "16:20",
                time_end: "17:20",
                title: "Лекция по React",
                location: "ул. V, 9АК, этаж 2",
                coords: [12.0, 12.0],
                tags: ["IT", "frontend"],
                external_url: "https://kudago.com/event/123123",
            },
            {
                type: "remote",
                time_start: "17:00",
                time_end: "18:00",
                title: "Математический анализ",
            },
            {
                type: "event",
                time_start: "16:20",
                time_end: "17:20",
                title: "Лекция по React",
                location: "ул. V, 9АК, этаж 2",
                coords: [12.0, 12.0],
                tags: ["IT", "frontend"],
                external_url: "https://kudago.com/event/123123",
            },
        ],
    },
];

export const globalMockData = {
    schedule: schedule,
};
