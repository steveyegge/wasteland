import { useEffect, useState } from "react";
import { toast } from "sonner";
import { profiles } from "../api/client";
import type { ProfileSummary } from "../api/types";
import { EmptyState } from "./EmptyState";
import styles from "./ProfilesList.module.css";
import { SkeletonRows } from "./Skeleton";

export function ProfilesList() {
  const [items, setItems] = useState<ProfileSummary[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  useEffect(() => {
    (async () => {
      try {
        const resp = await profiles();
        setItems(resp.profiles);
      } catch (e) {
        const msg = e instanceof Error ? e.message : "Failed to load";
        setError(msg);
        toast.error(msg);
      } finally {
        setLoading(false);
      }
    })();
  }, []);

  return (
    <div className={styles.page}>
      <h2 className={styles.heading}>Profiles</h2>

      {error && <p className={styles.error}>{error}</p>}

      {loading ? (
        <SkeletonRows count={6} />
      ) : items.length === 0 ? (
        <EmptyState
          title="No profiles yet"
          description="Profiles appear when rigs post, claim, or complete work on the wanted board."
        />
      ) : (
        <>
          <table className={styles.table} aria-label="Profiles">
            <thead>
              <tr className={styles.thead}>
                <th className={styles.th}>Handle</th>
                <th className={styles.thNum}>Posted</th>
                <th className={styles.thNum}>Claimed</th>
                <th className={styles.thNum}>Completed</th>
                <th className={styles.thNum}>Stamps</th>
              </tr>
            </thead>
            <tbody>
              {items.map((p) => (
                <tr key={p.handle} className={styles.row}>
                  <td className={styles.td}>
                    <span className={styles.handle}>{p.handle}</span>
                  </td>
                  <td className={styles.tdNum}>{p.posted}</td>
                  <td className={styles.tdNum}>{p.claimed}</td>
                  <td className={styles.tdNum}>{p.completed}</td>
                  <td className={styles.tdNum}>{p.stamps}</td>
                </tr>
              ))}
            </tbody>
          </table>

          <div className={styles.cardList}>
            {items.map((p) => (
              <div key={p.handle} className={styles.card}>
                <div className={styles.cardHandle}>{p.handle}</div>
                <div className={styles.cardStats}>
                  <span>{p.posted} posted</span>
                  <span>{p.claimed} claimed</span>
                  <span>{p.completed} completed</span>
                  <span>{p.stamps} stamps</span>
                </div>
              </div>
            ))}
          </div>
        </>
      )}
    </div>
  );
}
