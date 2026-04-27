import { useEffect, useState } from "react";

type Profile = {
  id: string;
  displayName: string;
};

async function fetchProfile(userId: string): Promise<Profile> {
  const response = await fetch(`/api/users/${encodeURIComponent(userId)}`);
  return response.json() as Promise<Profile>;
}

export function ProfilePanel({ userId }: { userId: string }) {
  const [profile, setProfile] = useState<Profile | null>(null);

  useEffect(() => {
    let cancelled = false;

    async function loadProfile(): Promise<void> {
      try {
        const nextProfile = await fetchProfile(userId);
        if (!cancelled) {
          setProfile(nextProfile);
        }
      } catch (error) {
        if (!cancelled) {
          console.error("profile load failed", error);
        }
      }
    }

    void loadProfile();

    return () => {
      cancelled = true;
    };
  }, [userId]);

  return <h2>{profile?.displayName ?? "Loading"}</h2>;
}
