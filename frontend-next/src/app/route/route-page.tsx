"use client";
import RouteList from "@/app/route/route-list";
import { buildRoute } from "@/lib/utils/build-yandex-maps-url";
import { TimelineItem } from "@/lib/api";
import Button from "@/components/button";

interface RoutePageProps {
    children: TimelineItem[];
}

export default function RoutePage({ children }: RoutePageProps) {
    const isRoutesEmpty = children.length === 0;

    const openMap = () => {
        if (isRoutesEmpty) return;

        const mapsUrl = `https://yandex.ru/maps/?rtext=${buildRoute(children)}`;
        window.open(mapsUrl, "_blank");
    };

    return (
        <div className="flex flex-col items-center justify-center h-screen">
            <h2 className="text-2xl font-bold mb-4">Маршрут на сегодня</h2>
            <div className="max-h-[80vh] overflow-y-auto">
                <RouteList>{children}</RouteList>
            </div>
            {isRoutesEmpty ? (
                <div className="mt-4 text-center text-sm text-gray-500">
                    Сегодня планов нет
                </div>
            ) : (
                <>
                    <Button onClick={openMap} className="cursor-pointer">
                        Открыть маршрут на карте
                    </Button>

                    <div className="text-[11px] text-(--text-3) text-center mt-1.5">
                        Откроется Яндекс.Карты с маршрутом
                    </div>
                </>
            )}
        </div>
    );
}
