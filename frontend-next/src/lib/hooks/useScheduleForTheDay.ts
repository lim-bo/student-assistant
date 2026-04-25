import { useQuery } from "@tanstack/react-query";
import { globalMockData } from "@/lib/api";

export const useScheduleForTheDay = (day: Date, group: number) => {
    return useQuery({
        queryKey: ["schedule"],
        queryFn: async () => {
            return globalMockData.schedule.find((item) => {
                return (
                    item.date.getDay() === day.getDay() && item.group === group
                );
            });
        },
    });
};
