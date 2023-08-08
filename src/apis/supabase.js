import { supabase } from 'src/supabase'

export const getCountries = async () => {
  const results = await supabase.from('countries').select()
  return results.data
}

export const signInMagicLinkStyle = async (email) => {
  const result = await supabase.auth.signInWithOtp({
    email,
    options: {
      emailRedirectTo: `${process.env.BASE_URL}/account`,
    },
  })
  return result
}

export const signOutUser = async () => {
  await supabase.auth.signOut()
}

export const getUserProfile = async (id) => {
  const result = await supabase
    .from('profiles')
    .select(`username, full_name, website, avatar_url`)
    .eq('id', id)
    .single()
  return result
}

export const updateUserProfile = async (updates) => {
  const result = await supabase.from('profiles').upsert(updates)
  return result
}

export const downloadAvatar = async (filePath) => {
  const result = await supabase.storage.from('avatars').download(filePath)
  return result
}

export const uploadAvatar = async (filePath, file) => {
  const result = await supabase.storage.from('avatars').upload(filePath, file)
  return result
}
