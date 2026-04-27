type User = {
  id: string;
  email: string;
};

async function sendWelcomeEmail(user: User): Promise<void> {
  await fetch(`/api/users/${encodeURIComponent(user.id)}/welcome`, {
    method: "POST",
  });
}

export async function inviteUsers(users: User[]): Promise<void> {
  try {
    for (const user of users) {
      await sendWelcomeEmail(user);
    }
  } catch (error) {
    console.error("failed to invite users", error);
    throw error;
  }
}
