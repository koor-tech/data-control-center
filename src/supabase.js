import { createClient } from '@supabase/supabase-js'

const supabaseUrl = ENV.SUPABASE_URL || process.env.SUPABASE_URL
const supabaseAnonKey = ENV.SUPABASE_ANON_KEY || process.env.SUPABASE_ANON_KEY

export const supabase = createClient(supabaseUrl, supabaseAnonKey)
