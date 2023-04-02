export const convertDate = (date?: string | undefined): string | undefined => 
  date ? new Date(date).toISOString() : undefined


export const parseDateToDatetimeLocal = (date: Date): string => {
  const strs = date.toISOString().split('T')
  return strs[0] + "T" + strs[1].slice(0, 5)
}

export const createQueryObject = (obj: any): string | undefined => {
  const query: any = {}
  for (const key in obj) {
    if (obj[key] !== undefined && obj[key] !== "") query[key] = obj[key] 
  }
  if(Object.keys(query).length === 0){
    return undefined
  }
  return query
}