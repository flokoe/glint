# Frontend Migration Plan: Go Templates → SvelteKit + TailwindCSS

## Overview
This document outlines the migration plan for converting the Clairvoyance AI Chat UI frontend from Go HTML templates with AlpineJS/HTMX to a modern SvelteKit application with TailwindCSS.

## Current Architecture Analysis

### Existing Components Structure
- **Base Template**: `00-base.html` - Main layout with background orbs
- **Chat View**: `01-chat-view.html` - Container with sidebar and main content
- **Chat Content**: `chat-content.html` - Message area and input form
- **Sidebar**: `sidebar.html` - Navigation with pinned/recent chats
- **Search Dialog**: `search-dialog.html` - Modal search interface
- **Admin Interface**: `admin.html` - Provider management

### Current Tech Stack
- **Frontend**: AlpineJS (3.14.9) + HTMX (2.0.6)
- **CSS**: Custom CSS with CSS variables, glassmorphism design
- **Templates**: Go HTML templates with server-side rendering
- **Build**: No build process, direct file serving

### API Endpoints Identified
- `GET /` - Chat view (renders template)
- `GET /api/chat` - Chat content (HTMX partial)
- `POST /api/completion` - Send message and get AI response
- `GET /admin` - Admin view
- `POST /api/admin/providers` - Add LLM provider

## Migration Strategy

### Phase 1: Project Setup & Configuration
1. **Initialize SvelteKit Project**
   - Create new SvelteKit project in `frontend/` directory
   - Configure for client-side rendering (CSR) mode
   - Set up TailwindCSS with custom color palette
   - Install additional dependencies (lucide-svelte for icons)

2. **Development Environment**
   - Configure Vite for development server
   - Set up proxy for backend API calls
   - Configure TailwindCSS with existing color scheme
   - Set up TypeScript for type safety

### Phase 2: Core Component Migration
1. **Layout Components**
   - `src/lib/components/Layout.svelte` - Main app layout
   - `src/lib/components/BackgroundOrbs.svelte` - Animated background
   - `src/lib/components/Sidebar.svelte` - Navigation sidebar
   - `src/lib/components/SidebarCollapsed.svelte` - Collapsed state

2. **Chat Components**
   - `src/lib/components/ChatArea.svelte` - Message display area
   - `src/lib/components/MessageInput.svelte` - Input form with model selection
   - `src/lib/components/Message.svelte` - Individual message component
   - `src/lib/components/LoadingIndicator.svelte` - AI response loading

3. **UI Components**
   - `src/lib/components/SearchDialog.svelte` - Search modal
   - `src/lib/components/ModelSelector.svelte` - LLM model dropdown
   - `src/lib/components/ChatListItem.svelte` - Chat list items

### Phase 3: State Management
1. **Stores Setup**
   - `src/lib/stores/chat.ts` - Chat state (messages, conversations)
   - `src/lib/stores/ui.ts` - UI state (sidebar, search, loading)
   - `src/lib/stores/models.ts` - Available LLM models
   - `src/lib/stores/settings.ts` - User preferences

2. **API Integration**
   - `src/lib/api/chat.ts` - Chat API calls
   - `src/lib/api/admin.ts` - Admin API calls
   - `src/lib/types/` - TypeScript interfaces

### Phase 4: Feature Implementation
1. **Chat Functionality**
   - Message sending and receiving
   - Real-time AI responses (Server-Sent Events or WebSocket)
   - Message history persistence
   - Model selection and configuration

2. **Navigation Features**
   - Sidebar toggle (Ctrl+B)
   - Search functionality (Ctrl+K, /)
   - New chat creation (Ctrl+Shift+O)
   - Chat history management

3. **Admin Features**
   - Provider management interface
   - Model configuration
   - Settings management

### Phase 5: Styling & Polish
1. **TailwindCSS Migration**
   - Convert CSS variables to Tailwind config
   - Implement glassmorphism effects with Tailwind utilities
   - Responsive design improvements
   - Dark theme optimization

2. **Animations & Interactions**
   - Smooth transitions between states
   - Loading animations
   - Hover effects and micro-interactions
   - Keyboard navigation enhancements

## Technical Implementation Details

### SvelteKit Configuration
```javascript
// vite.config.js
export default {
  plugins: [sveltekit()],
  server: {
    proxy: {
      '/api': 'http://localhost:1323'
    }
  }
}
```

### TailwindCSS Custom Configuration
```javascript
// tailwind.config.js
module.exports = {
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
        glass: {
          white: 'rgba(255, 255, 255, 0.1)',
          light: 'rgba(255, 255, 255, 0.05)',
          'ultra-light': 'rgba(255, 255, 255, 0.02)',
          border: 'rgba(255, 255, 255, 0.15)'
        },
        text: {
          primary: 'rgba(255, 255, 255, 0.95)',
          secondary: 'rgba(255, 255, 255, 0.7)',
          tertiary: 'rgba(255, 255, 255, 0.5)'
        }
      }
    }
  }
}
```

### Project Structure
```
frontend/
├── src/
│   ├── app.html
│   ├── app.css
│   ├── routes/
│   │   ├── +layout.svelte
│   │   ├── +page.svelte
│   │   └── admin/
│   │       └── +page.svelte
│   └── lib/
│       ├── components/
│       │   ├── Layout.svelte
│       │   ├── Sidebar.svelte
│       │   ├── ChatArea.svelte
│       │   ├── MessageInput.svelte
│       │   ├── SearchDialog.svelte
│       │   └── BackgroundOrbs.svelte
│       ├── stores/
│       │   ├── chat.ts
│       │   ├── ui.ts
│       │   └── models.ts
│       ├── api/
│       │   ├── chat.ts
│       │   └── admin.ts
│       └── types/
│           └── index.ts
├── static/
├── package.json
├── vite.config.js
├── tailwind.config.js
└── tsconfig.json
```

## Migration Steps

### Step 1: Setup New SvelteKit Project
1. Create `frontend/` directory
2. Initialize SvelteKit with TypeScript
3. Install dependencies (TailwindCSS, lucide-svelte)
4. Configure build tools and proxy

### Step 2: Migrate Layout & Navigation
1. Create main layout component
2. Implement sidebar with chat history
3. Add search dialog functionality
4. Set up keyboard shortcuts

### Step 3: Implement Chat Interface
1. Create chat area component
2. Build message input with model selection
3. Implement message display
4. Add loading states

### Step 4: API Integration
1. Set up API client functions
2. Implement chat message sending
3. Handle real-time responses
4. Add error handling

### Step 5: Admin Interface
1. Create admin page layout
2. Implement provider management
3. Add model configuration
4. Set up settings interface

### Step 6: Testing & Optimization
1. Test all functionality
2. Optimize performance
3. Ensure responsive design
4. Add accessibility features

## Benefits of Migration

### Developer Experience
- Modern component-based architecture
- TypeScript for type safety
- Hot module replacement during development
- Better debugging tools

### Performance
- Smaller bundle sizes with tree-shaking
- Client-side routing for faster navigation
- Optimized rendering with Svelte compilation
- Better caching strategies

### Maintainability
- Modular component structure
- Clear separation of concerns
- Reactive state management
- Reusable component library

### User Experience
- Smoother interactions
- Better accessibility
- Improved mobile responsiveness
- Enhanced keyboard navigation

## Considerations & Challenges

### Backend Integration
- Ensure API compatibility
- Handle CORS if needed
- Maintain existing authentication
- Consider WebSocket for real-time updates

### State Management
- Persist chat history locally
- Sync with backend state
- Handle offline scenarios
- Manage loading states

### Browser Compatibility
- Ensure compatibility with target browsers
- Progressive enhancement
- Fallback for older browsers
- Testing across devices

## Timeline Estimation

- **Phase 1**: 1-2 days (Setup & Configuration)
- **Phase 2**: 3-4 days (Core Component Migration)
- **Phase 3**: 2-3 days (State Management)
- **Phase 4**: 4-5 days (Feature Implementation)
- **Phase 5**: 2-3 days (Styling & Polish)

**Total Estimated Time**: 12-17 days

## Success Criteria

1. **Functional Parity**: All existing features work identically
2. **Performance**: Improved loading times and interactions
3. **Accessibility**: Better keyboard navigation and screen reader support
4. **Responsiveness**: Improved mobile and tablet experience
5. **Maintainability**: Cleaner, more modular codebase
6. **Type Safety**: Full TypeScript coverage

## Next Steps

1. Review and approve this migration plan
2. Set up development environment
3. Begin Phase 1 implementation
4. Establish testing procedures
5. Plan deployment strategy

---

*This migration plan provides a comprehensive roadmap for converting the Clairvoyance frontend to a modern SvelteKit application while maintaining all existing functionality and improving the overall user experience.*