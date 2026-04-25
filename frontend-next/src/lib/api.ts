const schedule = [
    {
        group: 3376,
        date: new Date(),
        timeline: [
            {
                type: "class",
                time_start: "11:40",
                time_end: "13:10",
                title: "Социология",
                location: "ул. Ломоносова, 9, ауд. 101",
                coords: [59.9398, 30.3178],
            },
            {
                type: "class",
                time_start: "13:40",
                time_end: "15:10",
                title: "Математический анализ",
                location: "ул. Ломоносова, 9, ауд. 106",
                coords: [59.9398, 30.3178],
            },
            {
                type: "event",
                time_start: "16:00",
                time_end: "17:00",
                title: "Лекция по Python",
                location: "ул. Пушкина, 9АК, этаж 2",
                coords: [59.9398, 30.3178],
                tags: ["IT", "backend"],
                external_url: "https://kudago.com/event/123",
            },
            {
                type: "class",
                time_start: "17:20",
                time_end: "18:50",
                title: "Математический анализ",
                location: "ул. Ломоносова, 9, ауд. 153",
                coords: [59.9398, 30.3178],
            },
        ],
    },
];

export const globalMockData = {
    schedule: schedule,
};
