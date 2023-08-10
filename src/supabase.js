import { createClient } from '@supabase/supabase-js'

const supabaseUrl = 'https://hobculzzveckkkreugyr.supabase.co'
const supabaseAnonKey =
  'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImhvYmN1bHp6dmVja2trcmV1Z3lyIiwicm9sZSI6ImFub24iLCJpYXQiOjE2OTA0MDUxNTAsImV4cCI6MjAwNTk4MTE1MH0.46NlkXzPfxCdxiQeVrZ2J6Lf0VSklBCbRe3RLnCsMUQ'

export const supabase = createClient(supabaseUrl, supabaseAnonKey)
