char* isValidOpen(char open, char* s);

char* isValidClose(char open, char* s) {
  if (*s == '\0') return NULL;
  char ch = *s;
  char* sub = s;
  while (ch != open && sub != NULL) {
    sub = isValidOpen(ch, sub+1);
    if (sub != NULL)
     ch = *sub;
  }
  if (ch == open) {
    sub = sub + 1;
  }
  return sub;
}

char* isValidOpen(char open, char* s) {
  switch (open){
    case '(': return isValidClose(')', s); 
    case '[': return isValidClose(']', s); 
    case '{': return isValidClose('}', s); 
    default: return NULL; 
  }
}


bool isValid(char * s){
  if (s == NULL) return false;
  if (*s == '\0') return true; 
  char* sub = isValidOpen(*s, s+1);
  return isValid(sub);
}