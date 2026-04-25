import {CardLayout} from "@/app/route/CardLayout";
import {InPersonClass, RemoteClass} from "@/lib/api";


export default function ClassCard({item}: { item: InPersonClass | RemoteClass }) {
  const isRemote = item.type === "remote";

  return (
    <CardLayout title={item.title} badge="занятие">
      <div className="text-sm text-gray-600">
        {item.time_start} – {item.time_end}
      </div>

      <div className="text-sm mt-2">
        {isRemote ? (
            <span>Онлайн</span>
          ) :
          (
            <span>📍 {item.location}</span>
          )
        }
      </div>
    </CardLayout>
  );
}