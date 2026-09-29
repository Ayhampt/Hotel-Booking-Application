# Notification Service

The Notification Service is an asynchronous, event-driven microservice responsible for formatting and dispatching transactional communications (emails) to users.

By fully decoupling email delivery from the core user-facing APIs via **Redis Message Queues**, this service ensures that network latency from SMTP providers does not impact the performance or response times of the `AuthInGo` or `BookingService` gateways.

## 🏗 Key Architectural Decisions

- **Event-Driven Design:** Operates as a background consumer that listens to specific Redis queues for job payloads.
- **Dynamic Templating:** Uses **Handlebars (`.hbs`)** to inject dynamic variables (like user names, booking dates, and OTPs) into pre-designed, responsive HTML email templates.
- **Stateless Processing:** Functions as a pure processor—fetching jobs from the queue, rendering the template, dispatching the email via SMTP, and acknowledging the job completion.

---

## 🔄 Event Architecture & Flow

```mermaid
graph TD
    %% Producers
    Auth["AuthInGo Service<br/>(User Signup)"] -->|Enqueues Payload| Redis[("Redis Queue")]
    Booking["Booking Service<br/>(Booking Confirmed)"] -->|Enqueues Payload| Redis

    %% Consumer
    Redis -->|Consumed By| Worker["Notification Service<br/>Queue Listener"]

    %% Processing
    Worker -->|Injects Data| HBS["Handlebars<br/>Template Engine"]
    HBS -->|HTML Output| Mailer["Email Provider<br/>Nodemailer/SMTP"]

    %% Output
    Mailer -.->|Delivers| User(["End User Inbox"])