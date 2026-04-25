import {CardLayout} from "@/app/route/CardLayout";
import {EventItem} from "@/lib/api";

export default function EventCard({item}: { item: EventItem }) {
  return (
    <CardLayout title={item.title} badge="мероприятие">
      <div className="text-sm text-gray-600">
        {item.time_start} – {item.time_end}
      </div>

      <div className="text-sm mt-2">
        📍 {item.location}
      </div>
    </CardLayout>
  );
}