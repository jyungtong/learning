import { GreetForm } from "@/components/greet-form";

export default function Home() {
  return (
    <main className="greet-page">
      <section aria-labelledby="greeting-title" className="greet-panel">
        <p className="service-name">greet.v1</p>
        <h1 id="greeting-title">Hello</h1>
        <GreetForm />
      </section>
    </main>
  );
}
