#ifdef __cplusplus
extern "C" {
#endif

// Phone info related
struct phone_info {
  char* number;
  char* normalized;
  char* error;
  int valid;
};
struct phone_info* new_phone_info(char*);
void free_phone_info(struct phone_info*);

// helpers
char* allocAndCopyStr(const char* src);

// Public api
int is_possible_number(char* number, char* region);
struct phone_info* parse(char* number, char* region);
char* get_country_code(char* number);

#ifdef __cplusplus
}
#endif
