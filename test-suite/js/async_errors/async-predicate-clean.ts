type Customer = {
  id: string;
  email: string;
};

async function hasPaidInvoice(customer: Customer): Promise<boolean> {
  const response = await fetch(`/api/customers/${encodeURIComponent(customer.id)}/paid`);
  return response.ok;
}

export async function paidCustomers(customers: Customer[]): Promise<Customer[]> {
  try {
    const paid: Customer[] = [];
    for (const customer of customers) {
      if (await hasPaidInvoice(customer)) {
        paid.push(customer);
      }
    }
    return paid;
  } catch (error) {
    console.error("failed to filter paid customers", error);
    throw error;
  }
}
