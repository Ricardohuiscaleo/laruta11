# La Ruta 11 - Development Guidelines

## Code Quality Standards

### JavaScript/React Conventions
- **Functional Components**: Use React functional components with hooks exclusively
- **State Management**: Leverage useState, useEffect, useMemo, useCallback for state and side effects
- **Component Structure**: Organize components with state declarations first, followed by effects, then handlers, and finally render logic
- **Naming**: Use camelCase for variables/functions, PascalCase for components
- **Props Destructuring**: Destructure props in function parameters for clarity
- **Event Handlers**: Prefix with `handle` (e.g., handleAddToCart, handleSubmit)

### PHP API Conventions
- **Headers First**: Always set headers at the top (Content-Type, CORS, Cache-Control)
- **JSON Responses**: Return consistent JSON structure with `success`, `error`, and data fields
- **Error Handling**: Use try-catch blocks with descriptive error messages
- **Database Connections**: Create PDO/mysqli connections with error mode enabled
- **Input Validation**: Validate all inputs before processing
- **HTTP Methods**: Check request method and respond appropriately

### CSS/Styling Standards
- **Tailwind First**: Use Tailwind utility classes for styling
- **Responsive Design**: Mobile-first approach with sm:, md:, lg: breakpoints
- **Custom Styles**: Use inline `<style>` blocks for component-specific animations
- **Color Palette**: Consistent use of orange (#f97316), red (#dc2626), green (#10b981)
- **Spacing**: Use clamp() for responsive sizing: `clamp(minPx, vw, maxPx)`

## Semantic Patterns

### React Component Patterns

#### Modal Pattern
```jsx
const Modal = ({ isOpen, onClose, children }) => {
  if (!isOpen) return null;
  
  return (
    <div className="fixed inset-0 bg-black/50 z-50" onClick={onClose}>
      <div className="bg-white rounded-xl" onClick={(e) => e.stopPropagation()}>
        {children}
      </div>
    </div>
  );
};
```

#### Form Handling Pattern
```jsx
const [formData, setFormData] = useState({
  field1: '',
  field2: ''
});

const handleSubmit = async (e) => {
  e.preventDefault();
  setLoading(true);
  try {
    const response = await fetch('/api/endpoint.php', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(formData)
    });
    const data = await response.json();
    if (data.success) {
      // Handle success
    }
  } catch (error) {
    console.error('Error:', error);
  } finally {
    setLoading(false);
  }
};
```

#### State Update Pattern
```jsx
// Immutable updates
setCart(prevCart => [...prevCart, newItem]);
setFormData(prev => ({...prev, field: value}));
```

### PHP API Patterns

#### Standard API Structure
```php
<?php
header('Content-Type: application/json');
header('Access-Control-Allow-Origin: *');
require_once '../config.php';

try {
    $conn = new mysqli(DB_SERVER, DB_USERNAME, DB_PASSWORD, DB_NAME);
    
    if ($conn->connect_error) {
        throw new Exception('Connection failed');
    }
    
    // Process request
    $data = json_decode(file_get_contents('php://input'), true);
    
    // Validate input
    if (!isset($data['required_field'])) {
        throw new Exception('Missing required field');
    }
    
    // Execute query
    $stmt = $conn->prepare("SELECT * FROM table WHERE id = ?");
    $stmt->bind_param("i", $data['id']);
    $stmt->execute();
    $result = $stmt->get_result();
    
    echo json_encode([
        'success' => true,
        'data' => $result->fetch_all(MYSQLI_ASSOC)
    ]);
    
} catch (Exception $e) {
    echo json_encode([
        'success' => false,
        'error' => $e->getMessage()
    ]);
}
?>
```

#### Database Query Pattern
```php
// Prepared statements always
$stmt = $conn->prepare("INSERT INTO table (col1, col2) VALUES (?, ?)");
$stmt->bind_param("ss", $value1, $value2);
$stmt->execute();

// Fetch results
$result = $stmt->get_result();
$data = $result->fetch_all(MYSQLI_ASSOC);
```

### File Upload Pattern (S3)
```php
// Use S3Manager class
$s3 = new S3Manager($config);
$imageUrl = $s3->uploadFile($_FILES['image'], "folder/filename.jpg", true);
```

## Architecture Patterns

### Component Organization
- **Atomic Design**: Break UI into small, reusable components
- **Container/Presentational**: Separate logic (containers) from UI (presentational)
- **Custom Hooks**: Extract reusable logic into custom hooks (e.g., useDoubleTap)

### API Design
- **RESTful Endpoints**: Use clear, resource-based URLs
- **Consistent Responses**: Always return JSON with success/error structure
- **Error Codes**: Use appropriate HTTP status codes (200, 400, 500)
- **Caching**: Set Cache-Control headers appropriately

### State Management
- **Local State**: Use useState for component-specific state
- **Derived State**: Use useMemo for computed values
- **Side Effects**: Use useEffect for data fetching and subscriptions
- **Refs**: Use useRef for DOM access and mutable values

## Common Code Idioms

### Conditional Rendering
```jsx
{isLoading && <Spinner />}
{error && <ErrorMessage message={error} />}
{data ? <DataDisplay data={data} /> : <EmptyState />}
```

### Array Operations
```jsx
// Filter
const filtered = items.filter(item => item.active);

// Map
const mapped = items.map(item => ({ ...item, newField: value }));

// Reduce
const total = items.reduce((sum, item) => sum + item.price, 0);
```

### Async/Await Pattern
```jsx
const fetchData = async () => {
  try {
    const response = await fetch('/api/endpoint.php');
    const data = await response.json();
    return data;
  } catch (error) {
    console.error('Error:', error);
    throw error;
  }
};
```

### Event Handling
```jsx
// Prevent default
const handleSubmit = (e) => {
  e.preventDefault();
  // Handle form
};

// Stop propagation
const handleClick = (e) => {
  e.stopPropagation();
  // Handle click
};
```

## Frequently Used Annotations

### JSDoc Comments
```jsx
/**
 * Adds item to cart with validation
 * @param {Object} product - Product to add
 * @param {number} quantity - Quantity to add
 * @returns {boolean} Success status
 */
const addToCart = (product, quantity) => {
  // Implementation
};
```

### PHP Comments
```php
// Single-line for brief explanations
/* Multi-line for complex logic */

/**
 * Uploads file to S3 with compression
 * @param array $file - Uploaded file array
 * @param string $key - S3 key/path
 * @return string - Public URL of uploaded file
 */
```

## Best Practices

### Performance
- Use useMemo for expensive calculations
- Use useCallback for event handlers passed to children
- Implement lazy loading for images and components
- Minimize re-renders with proper dependency arrays

### Security
- Sanitize all user inputs
- Use prepared statements for SQL queries
- Validate file uploads (type, size)
- Set appropriate CORS headers
- Never expose sensitive credentials in frontend

### Accessibility
- Use semantic HTML elements
- Add aria-labels to interactive elements
- Ensure keyboard navigation works
- Provide alt text for images

### Error Handling
- Always wrap async operations in try-catch
- Provide user-friendly error messages
- Log errors for debugging
- Implement fallback UI for errors

### Code Organization
- One component per file
- Group related files in folders
- Keep functions small and focused
- Extract reusable logic into utilities
- Use meaningful variable names
