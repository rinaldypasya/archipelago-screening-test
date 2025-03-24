# Basic Question

### **Designing the Photo Moderation System**

To build a website that allows users to submit photos and requires moderation before publishing, I would design the system with the following flow:

1.  **User Submission**
    
    -   A user uploads a photo via the frontend.
        
    -   The system stores the photo in a temporary storage location.
        
    -   Metadata (e.g., uploader, timestamp, AI analysis results) is saved in a database.
        
    -   The status of the photo is set to **"pending_review"**.
        
2.  **Automated & Manual Review**
    
    -   An AI-based moderation service (like AWS Rekognition, Google Vision, or an open-source model) scans the image for explicit content, violence, or prohibited content.
        
    -   If the AI detects issues, it flags the photo and marks it for manual review.
        
    -   If the AI does not detect issues, it moves to the **moderator review queue**.
        
3.  **Moderator Approval/Rejection**
    
    -   Moderators access a dashboard to review pending photos.
        
    -   They approve or reject the photo:
        
        -   **Approved:** Status changes to "published," and the image is moved to the public storage location.
            
        -   **Rejected:** Status changes to "rejected," and the user is notified.
            
        -   **Needs Revision:** User is asked to resubmit an adjusted image.
----------
### **Technology Stack**

#### **Frontend:**

-   **Vue.js** (for a reactive and user-friendly UI)
    
-   **Tailwind CSS** (for styling)
    
-   **Vue Router** (for navigation)
    
-   **Pinia/Vuex** (for state management)
    
-   **Axios** (for API communication)
    

#### **Backend:**

-   **Go (Golang)** (for handling API requests efficiently)
    
-   **Strapi (for admin panel and content management)**
    
-   **PostgreSQL** (for storing metadata about images, statuses, and reviews)
    
-   **Redis** (for caching pending reviews for quick access)
    
-   **MinIO / AWS S3** (for image storage)
    
-   **Celery with Redis or Go workers** (for async AI moderation)
    

#### **Infrastructure:**

-   **NGINX** (reverse proxy)
    
-   **PM2 / Systemd** (for process management)
    
-   **Docker + Kubernetes** (if scaling is needed)
    
----------

### **Data Structure in PostgreSQL**

I would define a table like this:

```
CREATE  TABLE photos (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID REFERENCES users(id) ON  DELETE CASCADE,
    image_url TEXT NOT  NULL,
    status TEXT CHECK (status IN ('pending_review', 'approved', 'rejected', 'needs_revision')) DEFAULT  'pending_review',
    ai_flags JSONB, -- Store AI analysis results created_at TIMESTAMP  DEFAULT NOW(),
    reviewed_by UUID REFERENCES moderators(id) NULL,
    reviewed_at TIMESTAMP  NULL,
    rejection_reason TEXT NULL );
```

### **Queue for Pending Reviews (Redis)**

-   Store `pending_review` photo IDs in a Redis list:
    ```
    LPUSH moderation_queue <photo_id>
    ```
    
-   When a moderator picks a photo:
    ```
    RPOP moderation_queue
    ```
    
This ensures quick retrieval of the next photo for moderation.


# Database Questions

## Level 1 (Novice)

Write a SQL query that shows me how many customers there are from Germany.
```
SELECT COUNT(CustomerID) FROM Customers WHERE Country = 'Germany';
```

## Level 2 (Business Admin)

Write a query that shows me a list of the countries that have the most customers; from most customers to least customers. Don’t show countries that have less than 5 customers.

```
SELECT COUNT(CustomerID) AS CustomerCount, Country
FROM Customers
GROUP BY Country
HAVING COUNT(CustomerID) >= 5
ORDER BY COUNT(CustomerID) DESC;
```

## Level 3 (Average Developer)

Reverse Engineer These Results (tell me the query that we need to write to get these results)
```
SELECT 
    Customers.CustomerName, 
    COUNT(Orders.OrderID) AS OrderCount, 
    MIN(Orders.OrderDate) AS FirstOrder, 
    MAX(Orders.OrderDate) AS LastOrder
FROM Customers, Orders
WHERE Customers.CustomerID = Orders.CustomerID
GROUP BY Customers.CustomerName
HAVING COUNT(Orders.OrderID) >= 5
ORDER BY MAX(Orders.OrderDate) DESC;
```

# Javascript/Typescript Questions

## Level 1: Title Case convertion.
Please go to `js/level1.js` file

## Level 2: Fix code using Promise.
Please go to `js/level2.js` file

## Level 2.5: Rewrite using Async/Await.
Please go to `js/level2half.js`

### **What’s Improved?**
- **Converted to Promises** → Eliminates the need for callbacks.  
- **Used `async/await`** → Makes the code more readable and easier to follow.  
- **Wrapped in `try/catch`** → Handles errors cleanly.  
- **No More Callback Hell** → Code is sequential and readable.

Now, the function `main()` runs asynchronously, fetching and processing data while handling errors properly! 🚀

## Level 3-4: Create a real-time chat between two windows (UNCOMPLETED YET)
Create a real-time chat between two windows; using web sockets, vuejs and typescript.


# Vue.Js

## Explain Vue.js reactivity and common issues when tracking changes.


### **Vue.js Reactivity (Simple Explanation)**

Vue.js **reactivity** means that when you change a piece of data, the UI **automatically updates** to reflect that change.

Vue tracks changes using a **reactive system** based on the **Proxy API (Vue 3)** or **Object.defineProperty (Vue 2)**. When a reactive property changes, Vue **detects the change** and **re-renders the affected parts of the UI**.

### **Common Issues When Tracking Changes**

#### **1. Not Using `ref()` or `reactive()` Properly (Vue 3)**

❌ **Wrong:**

```
const message = "Hello"; // NOT reactive` 
```

✅ **Correct:**

```
import { ref } from 'vue';
const message = ref("Hello"); // Now it's reactive` 
```

#### **2. Updating Objects or Arrays Incorrectly**

Vue **does not detect changes** when modifying objects/arrays **without replacing them**.

❌ **Wrong (Vue won't detect this change):**

```
const user = reactive({ name: "John" });
user.name = "Doe"; // Works, but Vue might not detect deeply nested changes
``` 

✅ **Correct:**

```
user = { ...user, name: "Doe" }; // Replace the object (for Vue 2) OR
user.name = "Doe"; // Works fine in Vue 3 (reactive tracks nested changes)
```

For **arrays**, use `.push()` or **create a new array**:

```
users.value = [...users.value, { name: "Alice" }]; // Works well
```

#### **3. Forgetting `.value` in Vue 3**

When using `ref()`, you **must** use `.value` to access the reactive value.

✅ **Correct:**

```
console.log(count.value); // Logs: 0
```

## Describe data flow between components in a Vue.js app

### Summary of Vue Data Flow

| Direction          | Method                   | Use Case                                    |
|--------------------|-------------------------|---------------------------------------------|
| **Parent → Child** | **Props**                | Pass data from parent to child             |
| **Child → Parent** | **Emit Events**          | Send data from child to parent             |
| **Sibling ↔ Sibling** | **State Management (Pinia/Vuex)** | Share data between unrelated components |
| **Global (Deep Components)** | **Provide/Inject** | Avoid prop drilling for deeply nested components |


## List the most common cause of memory leaks in Vue.js apps and how they can be solved.

### Summary of Vue.js Memory Leaks & Fixes

| Cause | Solution |
|-----------------------------|-----------------------------------------------------|
| **Unremoved Event Listeners** | Use `onBeforeUnmount()` to remove them. |
| **Unsubscribed Vuex/Pinia Watchers** | Store watcher in a variable and call `stop()`. |
| **Timers (setInterval, setTimeout) Not Cleared** | Use `clearInterval()` or `clearTimeout()`. |
| **DOM Element References Not Cleared** | Set references to `null` in `onBeforeUnmount()`. |
| **Global Event Bus (Vue 2)** | Use `$off()` to remove listeners. |
| **Third-Party Libraries Holding References** | Destroy or remove library instances in `onBeforeUnmount()`. |
| **Large Objects Stored in Reactive State** | Clear state manually when no longer needed. |
| **Too Many Active Vue Components** | Use `v-if` instead of `v-show` to fully remove components. |


## What have you used for state management

I've worked with various **state management solutions**, depending on the framework and use case. Here are some of the key ones:

-   **Pinia** (Recommended for Vue 3)
    
-   **Vuex** (Older but still used in Vue 2 apps)
    
-   **Composition API (reactive(), ref())** for local state management


## What’s the difference between pre-rendering and server side rendering?

Vue.js supports both **pre-rendering (SSG)** and **server-side rendering (SSR)**, primarily through **Nuxt.js**. Here’s how they differ:

| Feature | **Pre-Rendering (Static Site Generation - SSG)** | **Server-Side Rendering (SSR)** |
|---------|--------------------------------|--------------------------------|
| **Definition** | HTML is generated **at build time** and served as static files. | HTML is generated **on the server for each request** before being sent to the client. |
| **How it Works** | Uses **Nuxt generate** or **VuePress** to pre-build pages as static files. | Uses **Nuxt SSR mode** or **Vue SSR** to render pages dynamically on the server. |
| **Performance** | Very fast, served from CDN. | Slower than SSG, as each request triggers rendering. |
| **Use Case** | Best for static sites: blogs, documentation, landing pages. | Best for dynamic sites: dashboards, e-commerce, user-specific pages. |
| **SEO** | Excellent, as pages are pre-built and fully indexed. | Good SEO, but may have a slight delay compared to SSG. |
| **Data Freshness** | Static until the next build, needs regeneration for updates. | Always up-to-date, as pages are generated dynamically per request. |
| **Caching** | Can be fully cached on CDNs. | Can use server-side caching but still processes requests dynamically. |
| **Example in Vue/Nuxt** | `nuxt generate` (SSG mode), VuePress, VitePress | `nuxt start` (SSR mode), Vue + Express with SSR |

# Website Security Best Practises

- Enable HTTPS & Security Headers
- Implement Strong Authentication & MFA
- Sanitize User Input & Prevent SQL/XSS Attacks
- Keep Software & Dependencies Updated 
- Secure API Endpoints & Use Proper Authorization
- Monitor, Log, & Backup Data Regularly

# Website Performance Best Practises

- Optimize images & serve via CDN
- Minify & compress JavaScript, CSS, and HTML
- Reduce JavaScript execution time
- Enable caching & database optimizations 
- Reduce HTTP requests & third-party scripts
- Use lazy loading & code splitting
- Prefetch & preload important resources

# Golang: Counts the word frequency
Please go to `golang/main.go` file

# Tools (Rate yourself 1 to 5)

-   Git : 5
-   Redis : 5   
-   VSCode : 5
-   Linux : 4
-   AWS : 4
-   EC2 : 5
-   Lambda : 3
-   RDS : 5
-   Cloudwatch : 5
-   S3 : 5
-   Unit testing : 5
-   Kanban boards : 4