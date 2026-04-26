"use client";
import Card from "@/components/card";
import { useScheduleForTheDay } from "@/lib/hooks/useScheduleForTheDay";
import { buildRouteTimeline } from "@/lib/utils/build-route-timeline";
import { useSearchParams } from "next/navigation";
import { useMemo, useState } from "react";
import { twMerge as m } from "tailwind-merge";
import Button from "@/components/button";
import RoutePage from "@/app/route/route-page";

interface EventItem {
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

export default function Schedule() {
    const params = useSearchParams();
    const groupId = Number(params.get("group")) ?? 3376;
    const date = params.get("date") ?? "2010-10-10";
    const [routeStage, setRouteStage] = useState<boolean>(false);
    const [expandedDropdowns, setExpandedDropdowns] = useState<
        Record<string, boolean>
    >({});
    const [selectedEvents, setSelectedEvents] = useState<
        Record<string, EventItem>
    >({});
    const { data: schedule, isLoading } = useScheduleForTheDay(
        new Date(date),
        groupId,
    );

    const routeTimeline = useMemo(() => {
        if (!schedule) return [];
        return buildRouteTimeline(schedule, selectedEvents);
    }, [schedule, selectedEvents]);

    if (isLoading) {
        return <div>Loading...</div>;
    }

    const toggleDropdown = (classId: string) => {
        setExpandedDropdowns((prev) => ({
            ...prev,
            [classId]: !prev[classId],
        }));
    };

    const toggleEventSelection = (classId: string, event: EventItem) => {
        setSelectedEvents((prev) => {
            const currentSelectedForClass = prev[classId];
            if (currentSelectedForClass === event) {
                const newSelected = { ...prev };
                delete newSelected[classId];
                return newSelected;
            } else {
                return {
                    ...prev,
                    [classId]: event,
                };
            }
        });
    };

    const handlePostRoute = async () => {
        if (!schedule) return;
        setRouteStage(true);
    };

    return (
        <>
            {!routeStage ? (
                <div className="relative flex flex-col justify-center items-center min-h-screen">
                    <Button
                        className="fixed right-4 bottom-4 z-50"
                        onClick={handlePostRoute}
                        disabled={Object.keys(selectedEvents).length === 0}
                    >
                        <span>Рассчитать маршрут</span>
                    </Button>
                    <div className="relative items-center justify-center text-center inset-x-0">
                        <h2 className="text-2xl font-bold mb-4">
                            {new Date(date).toLocaleString("ru-RU", {
                                day: "numeric",
                                month: "long",
                            })}
                        </h2>
                        <div className="flex flex-col justify-center items-start w-full">
                            <div className="flex flex-col justify-center items-center gap-4">
                                {schedule?.map((classItem) => (
                                    <Card
                                        key={classItem.id}
                                        className="border rounded-xl p-3 bg-white shadow-sm border-black w-full"
                                    >
                                        <div className="flex w-full justify-between gap-6 items-start">
                                            <div
                                                className={m(
                                                    "flex flex-col font-mono tabular-nums w-fit",
                                                    classItem.type === "remote"
                                                        ? "text-blue-500"
                                                        : null,
                                                )}
                                            >
                                                <span className="text-4xl">
                                                    {classItem.time_start}
                                                </span>
                                                <span className="items-center justify-center text-center">
                                                    до
                                                </span>
                                                <span className="text-4xl">
                                                    {classItem.time_end}
                                                </span>
                                            </div>
                                            <div className="flex flex-col items-start justify-center flex-1">
                                                <span
                                                    className={m(
                                                        "font-medium",
                                                        classItem.type ===
                                                            "remote" &&
                                                            "font-thin italic",
                                                    )}
                                                >
                                                    {classItem.title}{" "}
                                                </span>
                                                <span
                                                    className={m(
                                                        "font-medium",
                                                        classItem.type ===
                                                            "remote" &&
                                                            "font-thin italic text-blue-500",
                                                    )}
                                                >
                                                    {classItem.type ===
                                                        "remote" &&
                                                        " (дистанционная)"}
                                                </span>
                                                <span className="text-gray-600">
                                                    {classItem.type ===
                                                        "class" &&
                                                        classItem.location}
                                                </span>
                                            </div>
                                            {classItem.type === "class" &&
                                                classItem.events.length > 0 && (
                                                    <button
                                                        onClick={() =>
                                                            toggleDropdown(
                                                                classItem.id,
                                                            )
                                                        }
                                                        className="flex items-center gap-2 px-3 py-2 rounded hover:bg-gray-100 transition-colors whitespace-nowrap"
                                                        aria-expanded={
                                                            expandedDropdowns[
                                                                classItem.id
                                                            ] || false
                                                        }
                                                    >
                                                        <span className="text-sm font-medium">
                                                            {
                                                                classItem.events
                                                                    .length
                                                            }{" "}
                                                            событие(й)
                                                        </span>
                                                        <span
                                                            className={`text-lg transition-transform inline-block ${
                                                                expandedDropdowns[
                                                                    classItem.id
                                                                ]
                                                                    ? "rotate-180"
                                                                    : ""
                                                            }`}
                                                        >
                                                            ▼
                                                        </span>
                                                    </button>
                                                )}
                                        </div>

                                        {classItem.type === "class" &&
                                            classItem.events.length > 0 &&
                                            expandedDropdowns[classItem.id] && (
                                                <div className="mt-4 pt-4 border-t border-gray-200 flex flex-col gap-3">
                                                    {classItem.events.map(
                                                        (event: EventItem) => (
                                                            <div
                                                                key={event.id}
                                                                className={m(
                                                                    "flex w-full justify-between gap-6 p-3 rounded transition-colors cursor-pointer border-2",
                                                                    selectedEvents[
                                                                        classItem
                                                                            .id
                                                                    ] === event
                                                                        ? "bg-green-100 shadow-lg shadow-green-300 border-green-400"
                                                                        : "bg-gray-50 hover:bg-gray-100 border-transparent",
                                                                )}
                                                                onClick={() =>
                                                                    toggleEventSelection(
                                                                        classItem.id,
                                                                        event,
                                                                    )
                                                                }
                                                            >
                                                                <div className="flex flex-col font-mono tabular-nums w-fit text-sm">
                                                                    <span>
                                                                        {
                                                                            event.time_start
                                                                        }
                                                                    </span>
                                                                    <span className="items-center justify-center text-center text-xs">
                                                                        до
                                                                    </span>
                                                                    <span>
                                                                        {
                                                                            event.time_end
                                                                        }
                                                                    </span>
                                                                </div>
                                                                <div className="flex flex-col items-start justify-center text-sm flex-1">
                                                                    <span className="font-medium">
                                                                        {
                                                                            event.title
                                                                        }
                                                                    </span>
                                                                    <span className="text-gray-600">
                                                                        {
                                                                            event.location
                                                                        }
                                                                    </span>
                                                                    {event.tags &&
                                                                        event
                                                                            .tags
                                                                            .length >
                                                                            0 && (
                                                                            <div className="flex gap-1 mt-1 flex-wrap">
                                                                                {event.tags.map(
                                                                                    (
                                                                                        tag: string,
                                                                                    ) => (
                                                                                        <span
                                                                                            key={
                                                                                                tag
                                                                                            }
                                                                                            className="text-xs bg-blue-100 text-blue-800 px-2 py-0.5 rounded"
                                                                                        >
                                                                                            {
                                                                                                tag
                                                                                            }
                                                                                        </span>
                                                                                    ),
                                                                                )}
                                                                            </div>
                                                                        )}
                                                                </div>
                                                            </div>
                                                        ),
                                                    )}
                                                </div>
                                            )}
                                    </Card>
                                ))}
                            </div>
                        </div>
                    </div>
                </div>
            ) : (
                schedule && <RoutePage>{routeTimeline}</RoutePage>
            )}
        </>
    );
}
