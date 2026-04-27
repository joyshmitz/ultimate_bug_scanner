type Customer = {
  id: string;
  email: string;
};

async function hasPaidInvoice(customer: Customer): Promise<boolean> {
  const response = await fetch(`/api/customers/${encodeURIComponent(customer.id)}/paid`);
  return response.ok;
}

export async function paidCustomers(customers: Customer[]): Promise<Customer[]> {
  return customers.filter(async (customer) => {
    return hasPaidInvoice(customer);
  });
}
