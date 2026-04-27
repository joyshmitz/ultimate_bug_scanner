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
  users.forEach(async (user) => {
    await sendWelcomeEmail(user);
  });
}
