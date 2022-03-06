#ifdef __cplusplus
 extern "C" {
#endif

typedef struct Parser {
	void *_parser;
} Parser;

int Create(Parser *p);
int AddBuffer(Parser *p, char *b);
unsigned int GetSize(Parser *p);
void Finish(Parser *p);
unsigned char *GetBuffer(Parser *p);
void Destroy(Parser *p);

int parseJSON(char *schema, char *json);
#ifdef __cplusplus
}
#endif

